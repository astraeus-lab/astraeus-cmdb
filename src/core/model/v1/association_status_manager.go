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

type AssociationStatusManager struct {
	dbClient    *gorm.DB
	cacheClient redis.UniversalClient
}

func NewAssociationStatusManager(options NewOptions) *AssociationStatusManager {
	res := &AssociationStatusManager{
		dbClient:    db.GetDBConnect(),
		cacheClient: cache.GetCacheConnect(),
	}

	if res.dbClient == nil {
		res.dbClient = db.GetDBConnect()
	}
	if res.cacheClient == nil {
		res.cacheClient = cache.GetCacheConnect()
	}

	return res
}

func (asm *AssociationStatusManager) List() (map[string]*model.AssociationStatus, error) {
	var queryAssociationStatus []db.ModelAssociationStatus
	err := asm.dbClient.
		Select("name", "association_status_uid", "description").
		Find(&queryAssociationStatus).Error
	if err != nil {
		return nil, err
	}

	res := make(map[string]*model.AssociationStatus)
	for idx := 0; idx < len(queryAssociationStatus); idx++ {
		res[queryAssociationStatus[idx].AssociationStatusUID] = &model.AssociationStatus{
			Name:                 queryAssociationStatus[idx].Name,
			AssociationStatusUID: queryAssociationStatus[idx].AssociationStatusUID,
			Description:          queryAssociationStatus[idx].Description,
		}
	}

	return res, nil
}

func (asm *AssociationStatusManager) Create(source *model.AssociationStatus) error {
	if asm.isExistAssociation(source.AssociationStatusUID) {
		return fmt.Errorf("%v err: association uid(%s) ", util.AlreadyExistErr, source.AssociationStatusUID)
	}

	return asm.dbClient.
		Create(&db.ModelAssociationStatus{
			Name:                 source.Name,
			AssociationStatusUID: source.AssociationStatusUID,
			Description:          source.Description,
		}).Error
}

func (asm *AssociationStatusManager) Update(associationStatusUID string, source *model.AssociationStatus) error {

	return asm.dbClient.
		Model(&db.ModelAssociationStatus{}).
		Where("association_status_uid = ?", associationStatusUID).
		Updates(&db.ModelAssociationStatus{
			Name:        source.Name,
			Description: source.Description,
		}).Error
}

func (asm *AssociationStatusManager) Retrivev(associationStatusUID string) (*model.AssociationStatus, error) {
	query := &db.ModelAssociationStatus{}
	err := asm.dbClient.
		Select("name", "association_status_uid", "description").
		Where("association_status_uid = ?", associationStatusUID).
		Take(query).Error
	if err != nil {
		return nil, err
	}

	return &model.AssociationStatus{
		Name:                 query.Name,
		Description:          query.Description,
		AssociationStatusUID: query.AssociationStatusUID,
	}, nil
}

func (asm *AssociationStatusManager) Delete(associationStatusUID string) error {
	if !asm.isExistAssociation(associationStatusUID) {
		return fmt.Errorf("%v err: association status uid(%s)", util.NotFoundErr, associationStatusUID)
	}

	var existAssociation []db.ModelAssociation
	err := asm.dbClient.
		Select("association_status_uid").
		Where("association_status_uid = ?", associationStatusUID).
		Find(&existAssociation).Error
	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("association(%s) still has related model or other err: %v", associationStatusUID, err)
	}

	return asm.dbClient.
		Delete(&db.ModelAssociationStatus{}).
		Where("association_status_uid = ?", associationStatusUID).Error
}

func (asm *AssociationStatusManager) isExistAssociation(associationStatusUID string) bool {
	err := asm.dbClient.
		Select("association_status_uid").
		Where("association_status_uid", associationStatusUID).
		Take(&db.ModelAssociationStatus{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
