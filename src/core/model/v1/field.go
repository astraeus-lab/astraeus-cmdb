package v1

import (
	"errors"
	"fmt"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/cache"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/db"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// TODO: The deletion and addition of fields require triggering a secondary update of the corresponding collector

type ModelField struct {
	modelUID    string
	dbClient    *gorm.DB
	cacheClient redis.UniversalClient
}

func NewModelField(modelUID string, options NewOptions) *ModelField {
	res := &ModelField{
		modelUID:    modelUID,
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

func (mf *ModelField) ListFiled() (map[string]*model.Field, error) {
	res := make(map[string]*model.Field)

	var queryField []db.ModelField
	err := mf.dbClient.
		Select("name", "field_uid", "is_editable", "is_optional", "description").
		Where("model_uid = ?", mf.modelUID).
		Find(&queryField).Error
	if err != nil {
		return nil, err
	}

	for idx := 0; idx < len(queryField); idx++ {
		res[queryField[idx].FieldUID] = &model.Field{
			Name:        queryField[idx].Name,
			FieldUID:    queryField[idx].FieldUID,
			IsEditable:  queryField[idx].IsEditable,
			IsOptional:  queryField[idx].IsOptional,
			Description: queryField[idx].Description,
		}
	}

	return res, nil
}

func (mf *ModelField) AddField(field *model.Field) error {
	if mf.isExistFieldUID(field.FieldUID) {
		return fmt.Errorf("%v err: field uid(%s)", util.AlreadyExistErr, field.FieldUID)
	}

	return mf.dbClient.
		Create(&db.ModelField{
			Name:          field.Name,
			FieldUID:      field.FieldUID,
			ModelUIDRefer: mf.modelUID,
			IsEditable:    field.IsEditable,
			IsOptional:    field.IsOptional,
			Description:   field.Description,
		}).Error
}

func (mf *ModelField) UpdateFiled(fieldUID string, field *model.Field) error {
	if !mf.isExistFieldUID(fieldUID) {
		return fmt.Errorf("%v err: field uid(%s)", util.AlreadyExistErr, field.FieldUID)
	}

	err := mf.dbClient.
		Model(&db.ModelField{}).
		Where("model_uid = ? and field_uid = ?", mf.modelUID, fieldUID).
		Updates(&db.ModelField{
			Name:        field.Name,
			IsEditable:  field.IsEditable,
			IsOptional:  field.IsOptional,
			Description: field.Description,
		}).Error

	return err
}

func (mf *ModelField) RetrieveField(fieldUID string) (*model.Field, error) {
	queryFiled := &db.ModelField{}
	err := mf.dbClient.
		Select("name", "field_uid", "is_editable", "description", "is_optional").
		Where("model_uid = ?", fieldUID).
		Take(queryFiled).Error
	if err != nil {
		return nil, err
	}

	return &model.Field{
		Name:        queryFiled.Name,
		FieldUID:    queryFiled.FieldUID,
		IsEditable:  queryFiled.IsEditable,
		IsOptional:  queryFiled.IsOptional,
		Description: queryFiled.Description,
	}, nil
}

func (mf *ModelField) DeleteField(fieldUID string) error {
	if !mf.isExistFieldUID(fieldUID) {
		return fmt.Errorf("%v err: field uid(%s)", util.NotFoundErr, fieldUID)
	}

	err := mf.dbClient.Transaction(func(tx *gorm.DB) error {
		deleteField, err := mf.RetrieveField(fieldUID)
		if err != nil {
			return err
		}

		if err = tx.
			Delete(deleteField).
			Where("field_uid", fieldUID).Error; err != nil {
			return err
		}

		if err = tx.
			Delete(&db.ModelResourceFieldData{}).
			Where("filed_uid", fieldUID).Error; err != nil {
			return err
		}

		return err
	})

	return err
}

func (mf *ModelField) isExistFieldUID(fieldUID string) bool {
	err := mf.dbClient.
		Select("field_uid").
		Where("field_uid = ?", fieldUID).
		Take(&db.ModelField{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
