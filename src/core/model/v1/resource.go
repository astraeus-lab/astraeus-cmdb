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

type ModelResource struct {
	modelUID    string
	dbClient    *gorm.DB
	cacheClient redis.UniversalClient
}

func NewModelResource(modelUID string, options NewOptions) *ModelResource {
	res := &ModelResource{
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

func (mr *ModelResource) ListResource() (map[string]*model.Resource, error) {
	res := make(map[string]*model.Resource)

	// TODO: try cache, if obtained from DB then add all to cache

	var allResource []db.ModelResource
	err := mr.dbClient.
		Select("resource_uid", "collection_uid", "is_auto_collected").
		Where("model_uid = ?", mr.modelUID).
		Find(&allResource).Error
	if err != nil {
		return nil, err
	}

	var allResourceData []db.ModelResourceFieldData
	err = mr.dbClient.
		Select("data", "field_uid", "resource_uid").
		Where("model_uid = ?", mr.modelUID).
		Find(allResourceData).Error
	if err != nil {
		return nil, err
	}

	for idx := 0; idx < len(allResource); idx++ {
		res[allResource[idx].ResourceUID] = &model.Resource{
			ResourceUID:       allResource[idx].ResourceUID,
			DataCollectionUID: allResource[idx].CollectionUIDRefer,
			IsAutoCollected:   allResource[idx].IsAutoCollected,
			FieldConenct:      make(map[string]string),
		}
	}

	if err != nil {
		return nil, err
	}
	for idx := 0; idx < len(allResourceData); idx++ {
		resrouceUID := allResourceData[idx].ResourceUIDRefer
		if res[resrouceUID] != nil {
			res[resrouceUID].FieldConenct[allResourceData[idx].FieldUIDRefer] = allResourceData[idx].Data
		}
	}

	return res, nil
}

func (mr *ModelResource) AddResource(resource *model.Resource) error {
	if mr.isExistResourceUID(resource.ResourceUID) {
		return fmt.Errorf("%v err: resource uid(%s) ", util.AlreadyExistErr, resource.ResourceUID)
	}

	addAllRecord := make([]db.ModelResourceFieldData, 0)
	for fieldUID, fieldContent := range resource.FieldConenct {
		addAllRecord = append(addAllRecord, db.ModelResourceFieldData{
			Data:             fieldContent,
			ModelUIDRefer:    mr.modelUID,
			FieldUIDRefer:    fieldUID,
			ResourceUIDRefer: resource.ResourceUID,
		})
	}

	err := mr.dbClient.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Create(&db.ModelResource{
				ModelUIDRefer:      mr.modelUID,
				ResourceUID:        resource.ResourceUID,
				CollectionUIDRefer: resource.DataCollectionUID,
				IsAutoCollected:    resource.IsAutoCollected,
			}).Error

		if err = tx.Create(&addAllRecord).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (mr *ModelResource) UpdateResource(resourceUID string, resource *model.Resource) error {
	oldResource, err := mr.RetrieveResource(resourceUID)
	if err != nil {
		return err
	}

	// TODO: update db first and then cache

	allFieldUID, err := mr.getAllFieldUIDByModelUID()
	if err != nil {
		return err
	}
	updateContent := make(map[string]string)
	for fieldUID := range allFieldUID {
		// When providing empty data, the update of the field content will be ignored,
		// and the deletion of the field content should be maintained
		// by the corresponding model field.
		if resource.FieldConenct[fieldUID] == "" ||
			resource.FieldConenct[fieldUID] == oldResource.FieldConenct[fieldUID] {

			continue
		}

		updateContent[fieldUID] = resource.FieldConenct[fieldUID]
	}
	err = mr.dbClient.Transaction(func(tx *gorm.DB) error {
		err = tx.
			Model(&db.ModelResource{}).
			Where("resource_uid = ?", resourceUID).
			Updates(&db.ModelResource{
				CollectionUIDRefer: resource.DataCollectionUID,
				IsAutoCollected:    resource.IsAutoCollected,
			}).Error
		if err != nil {
			return err
		}

		for fieldUID, content := range updateContent {
			err = tx.
				Model(&db.ModelResourceFieldData{}).
				Where("field_uid = ? AND resource_uid = ?", fieldUID, resource.ResourceUID).
				Update("data", content).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (mr *ModelResource) RetrieveResource(resourceUID string) (*model.Resource, error) {
	// TODO: try cache, if obtained from DB then add to cache

	var resourceData []db.ModelResourceFieldData
	err := mr.dbClient.
		Select("data", "field_uid").
		Where("resource_uid = ?", resourceUID).
		Find(&resourceData).Error
	if err != nil {
		return nil, err
	}

	allFieldUID, err := mr.getAllFieldUIDByModelUID()
	if err != nil {
		return nil, err
	}
	fieldContent := make(map[string]string)
	for idx := 0; idx < len(resourceData); idx++ {
		fieldUID := resourceData[idx].FieldUIDRefer
		if allFieldUID[fieldUID] {
			fieldContent[fieldUID] = resourceData[idx].Data
		}
	}

	resource := db.ModelResource{}
	err = mr.dbClient.
		Select("collection_uid", "is_auto_collected").
		Where("resource_uid = ? AND model_uid = ?", resourceUID, mr.modelUID).
		Take(&resource).Error
	if err != nil {
		return nil, err
	}

	return &model.Resource{
		ResourceUID:       resource.ResourceUID,
		DataCollectionUID: resource.CollectionUIDRefer,
		IsAutoCollected:   resource.IsAutoCollected,
		FieldConenct:      fieldContent,
	}, nil
}

func (mr *ModelResource) DeleteResource(resourceUID string) error {
	if !mr.isExistResourceUID(resourceUID) {
		return fmt.Errorf("%v err: resource uid(%s)", util.NotFoundErr, resourceUID)
	}

	// TODO: delete db first and then cache

	err := mr.dbClient.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Delete(&db.ModelResourceFieldData{}).
			Where("resource_id = ?", resourceUID).Error
		if err != nil {
			return err
		}

		err = tx.
			Delete(&db.ModelResource{}).
			Where("model_uid = ? AND resource_uid = ?", mr.modelUID, resourceUID).Error

		return err
	})

	return err
}

func (mr *ModelResource) isExistResourceUID(uid string) bool {
	err := mr.dbClient.
		Select("resource_uid").
		Where("resource_uid", uid).
		Take(&db.ModelAssociationStatus{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (mr *ModelResource) getAllFieldUIDByModelUID() (map[string]bool, error) {
	var fieldUID []string
	err := mr.dbClient.
		Model(&db.ModelField{}).
		Pluck("field_uid", &fieldUID).
		Where("model_uid = ?", mr.modelUID).Error
	if err != nil {
		return nil, err
	}

	res := make(map[string]bool)
	for idx := 0; idx < len(fieldUID); idx++ {
		res[fieldUID[idx]] = true
	}

	return res, nil
}
