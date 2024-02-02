package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/cache"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/db"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"

	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// ModelManager Astraeus-CMDB universal model.
//
// It is not safe for concurrent use by multiple
// goroutines without additional locking or coordination.
type ModelManager struct {
	dbClient    *gorm.DB
	cacheClient redis.UniversalClient
}

func NewModelManager(options NewOptions) *ModelManager {
	res := &ModelManager{
		dbClient:    options.DBClient,
		cacheClient: options.CacheClient,
	}

	if res.dbClient == nil {
		res.dbClient = db.GetDBConnect()
	}
	if res.cacheClient == nil {
		res.cacheClient = cache.GetCacheConnect()
	}

	return res
}

func (mm *ModelManager) List() (map[string]*model.Model, error) {
	var queryModel []db.Model
	err := mm.dbClient.
		Select("name", "revision", "model_uid", "group_uid", "is_builtin", "gmt_create", "gmt_modified").
		Find(&queryModel).Error
	if err != nil {
		return nil, err
	}

	res := make(map[string]*model.Model)
	for idx := 0; idx < len(queryModel); idx++ {
		res[queryModel[idx].ModelUID] = &model.Model{
			Metadata: model.Metadata{
				UID:            queryModel[idx].ModelUID,
				Name:           queryModel[idx].Name,
				GroupUID:       queryModel[idx].GroupUIDRefer,
				IsBuiltin:      queryModel[idx].IsBuiltin,
				CreateTime:     queryModel[idx].CreateAT.Unix(),
				LastUpdateTime: queryModel[idx].UpdateAT.Unix(),
				Revision:       strconv.Itoa(queryModel[idx].Revision),
			},
		}
	}

	return res, nil
}

func (mm *ModelManager) Create(metadata *model.Metadata) (err error) {
	if mm.isExistModelUID(metadata.UID) {
		return fmt.Errorf("%v err: model uid(%s)", util.AlreadyExistErr, metadata.UID)
	}

	revision := 0
	if revision, err = strconv.Atoi(metadata.Revision); err != nil {
		return fmt.Errorf("revision should be of integer, %v: %v", util.DataTypeErr, err)
	}
	if err = mm.dbClient.
		Create(&db.Model{
			Name:          metadata.Name,
			Revision:      revision,
			ModelUID:      metadata.UID,
			GroupUIDRefer: metadata.GroupUID,
			IsBuiltin:     false,
		}).Error; err != nil {
		return fmt.Errorf("create model err: %v", err)
	}

	return
}

func (mm *ModelManager) Update(modelUID string, metadata *model.Metadata) error {
	revision, err := strconv.Atoi(metadata.Revision)
	if err != nil {
		return fmt.Errorf("revision should be of integer, %v err: %v", util.DataTypeErr, err)
	}

	return mm.dbClient.
		Model(&db.Model{}).
		Where("model_uid = ?", modelUID).
		Updates(&db.Model{
			Name:     metadata.Name,
			Revision: revision,
		}).Error
}

func (mm *ModelManager) Retrieve(modelUID string, withData bool) (*model.Model, error) {
	res := &model.Model{}

	getModelMetada := func(dbClient *gorm.DB) error {
		queryModel := &db.Model{}
		err := dbClient.
			Select("name", "model_uid", "group_uid", "is_builtin", "gmt_create", "gmt_modified").
			Where("model_uid = ?", modelUID).
			Take(queryModel).Error
		if err != nil {
			return err
		}

		res = &model.Model{
			Metadata: model.Metadata{
				UID:            queryModel.ModelUID,
				Name:           queryModel.Name,
				Revision:       strconv.Itoa(queryModel.Revision),
				GroupUID:       queryModel.GroupUIDRefer,
				IsBuiltin:      queryModel.IsBuiltin,
				CreateTime:     queryModel.CreateAT.Unix(),
				LastUpdateTime: queryModel.UpdateAT.Unix(),
			},
		}

		return nil
	}

	if !withData {
		if err := getModelMetada(mm.dbClient); err != nil {
			return nil, err
		}
	} else {
		err := mm.dbClient.Transaction(func(tx *gorm.DB) error {
			if err := getModelMetada(tx); err != nil {
				return err
			}

			listFiled, err := NewModelField(modelUID, NewOptions{DBClient: tx}).ListFiled()
			if err != nil {
				return err
			}
			res.Field = listFiled

			listResource, err := NewModelResource(modelUID, NewOptions{DBClient: tx}).ListResource()
			if err != nil {
				return err
			}
			res.Resource = listResource

			// TODO: import associated data?

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (mm *ModelManager) Delete(modelUID string) error {
	if !mm.isExistModelUID(modelUID) {
		return fmt.Errorf("%v err: model uid(%s)", util.NotFoundErr, modelUID)
	}

	err := mm.dbClient.Transaction(func(tx *gorm.DB) error {
		deleteModel, err := mm.Retrieve(modelUID, false)
		if err != nil {
			return err
		}

		if err = tx.
			Delete(deleteModel).
			Where("model_uid", modelUID).Error; err != nil {
			return err
		}

		var allFieldUID []db.ModelField
		if err = tx.
			Select("field_uid").
			Where("model_uid", modelUID).
			Find(&allFieldUID).Error; err != nil {
			return err
		}

		for idx := 0; idx < len(allFieldUID); idx++ {
			if err = tx.
				Delete(&db.ModelResourceFieldData{}).
				Where("field_uid = ?", allFieldUID[idx].FieldUID).Error; err != nil {
				return err
			}
		}

		if err = tx.
			Delete(&db.ModelField{}).
			Where("model_uid = ?", modelUID).Error; err != nil {
			return err
		}

		if err = tx.
			Delete(&db.ModelResource{}).
			Where("model_uid = ?", modelUID).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (mm *ModelManager) SerializeExport(modelUID, exportType string, withData bool) (string, error) {
	source, err := mm.Retrieve(modelUID, withData)
	if err != nil {
		return "", err
	}

	marshalType := strings.ToUpper(exportType)
	switch marshalType {
	case "JSON":
		data, err := json.Marshal(source)
		if err != nil {
			return "", fmt.Errorf("%s serialization err: %v", marshalType, err)
		}
		return string(data), nil

	case "YAML":
		data, err := yaml.Marshal(source)
		if err != nil {
			return "", fmt.Errorf("%s serialization err: %v", marshalType, err)
		}
		return string(data), nil

	default:
		return "", fmt.Errorf("%s serialization type not supported", marshalType)
	}
}

func (mm *ModelManager) SerializeImport(source string) (*model.Model, error) {
	res := &model.Model{}
	data := []byte(source)

	if json.Valid(data) {
		if err := json.Unmarshal(data, res); err != nil {
			return nil, fmt.Errorf("JSON unmarshal err: %v", err)
		}

		return res, nil
	}

	if err := yaml.Unmarshal(data, res); err != nil {
		return nil, fmt.Errorf("YAML unmarshal err: %v", err)
	}

	// TODO: insert data

	return res, nil
}

func (mm *ModelManager) isExistModelUID(modelUID string) bool {

	// TODO: try cache

	err := mm.dbClient.
		Select([]string{"model_uid"}).
		Where("model_uid = ?", modelUID).
		Take(&db.Model{}).Error
	// Any query error is judged as UID already exists,
	// except for ErrRecordNotFound error.
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
