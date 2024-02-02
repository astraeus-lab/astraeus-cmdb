package v1

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/cache"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/db"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ModelAssociation struct {
	modelUID    string
	dbClient    *gorm.DB
	cacheClient redis.UniversalClient
}

func NewModelAssociation(modelUID string, options NewOptions) *ModelAssociation {
	res := &ModelAssociation{
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

func (ma *ModelAssociation) ListAssociation() (map[string]*model.Association, error) {
	res := make(map[string]*model.Association)

	var queryAssociation []db.ModelAssociation
	err := ma.dbClient.
		Select("association_uid", "model_uid", "association_model_uid", "association_status_uid").
		Where("model_uid = ?", ma.modelUID).
		Find(&queryAssociation).Error
	if err != nil {
		return nil, err
	}

	allAssociationStatus, err := NewAssociationStatusManager(NewOptions{DBClient: ma.dbClient}).List()
	if err != nil {
		return nil, err
	}

	allRelatedModel := make(map[string][]string)
	for idx := 0; idx < len(queryAssociation); idx++ {
		associationUID := queryAssociation[idx].AssociationStatusUIDRefer
		if allRelatedModel[associationUID] == nil {
			allRelatedModel[associationUID] = make([]string, 0)
		}
		allRelatedModel[associationUID] = append(allRelatedModel[associationUID], queryAssociation[idx].AssociationModelUIDRefer)
	}

	for associationUID, relatedModel := range allRelatedModel {
		res[associationUID] = &model.Association{
			Name:                 allAssociationStatus[associationUID].Name,
			Description:          allAssociationStatus[associationUID].Description,
			AssociationStatusUID: allAssociationStatus[associationUID].AssociationStatusUID,
			AssociationModelUID:  relatedModel,
		}
	}

	return res, err
}

func (ma *ModelAssociation) AddAssociation(relateModelUID, associationStatusUID string) error {
	if _, err := NewAssociationStatusManager(NewOptions{DBClient: ma.dbClient}).Retrivev(associationStatusUID); err != nil {
		return fmt.Errorf("%v err: association uid(%s)s", util.NotFoundErr, associationStatusUID)
	}

	associationUID := strings.Join([]string{ma.modelUID, associationStatusUID, relateModelUID}, "_")
	if ma.isExistAssociationUID(associationUID) {
		return fmt.Errorf("%v err: association uid(%s)", util.AlreadyExistErr, associationUID)
	}

	err := ma.dbClient.
		Create(&db.ModelAssociation{
			AssociationUID:            associationUID,
			ModelUIDRefer:             ma.modelUID,
			AssociationModelUIDRefer:  relateModelUID,
			AssociationStatusUIDRefer: associationStatusUID,
		}).Error

	return err
}

func (ma *ModelAssociation) RetrieveAssociation(associationStatusUID string) (*model.Association, error) {
	association := &db.ModelAssociationStatus{}
	err := ma.dbClient.
		Select("name", "association_status_uid", "description").
		Where("association_status_uid = ?", associationStatusUID).
		Take(&association).Error
	if err != nil {
		return nil, err
	}

	var allModelUID []string
	err = ma.dbClient.
		Model(&db.ModelAssociation{}).
		Pluck("association_model_uid", &allModelUID).
		Where("model_uid = ? AND association_status_uid = ?", ma.modelUID, associationStatusUID).Error
	if err != nil {
		return nil, err
	}

	return &model.Association{
		Name:                 association.Name,
		Description:          association.Description,
		AssociationStatusUID: association.AssociationStatusUID,
		AssociationModelUID:  allModelUID,
	}, nil
}

func (ma *ModelAssociation) DeleteAssociation(associationUID string) error {
	if !ma.isExistAssociationUID(associationUID) {
		return fmt.Errorf("%v err: association uid(%s)", util.NotFoundErr, associationUID)
	}

	return ma.dbClient.
		Delete(&db.ModelAssociation{}).
		Where("association_uid = ?", associationUID).Error
}

func (ma *ModelAssociation) isExistAssociationUID(associationUID string) bool {
	err := ma.dbClient.
		Select("association_uid").
		Where("association_uid = ? ", associationUID).
		Take(&db.ModelAssociation{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
