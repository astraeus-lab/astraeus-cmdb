package db

import (
	"time"
)

type ModelGroup struct {
	ID          uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	Name        string    `gorm:"type:VARCHAR(255) NOT NULL"`
	GroupUID    string    `gorm:"type:VARCHAR(63) NOT NULL;unique;uniqueIndex:uk_model_group_uid"`
	Description string    `gorm:"type:VARCHAR(1023)"`
	CreateAT    time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT    time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type Model struct {
	ID            uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	Name          string    `gorm:"type:VARCHAR(255) NOT NULL"`
	Revision      int       `gorm:"type:BIGINT UNSIGNED NOT NULL"`
	ModelUID      string    `gorm:"type:VARCHAR(63) NOT NULL;unique;uniqueIndex:uk_model_uid"`
	GroupUIDRefer string    `gorm:"type:VARCHAR(63) NOT NULL;column:group_uid"`
	IsBuiltin     bool      `gorm:"type:UNSIGNED TINYINT;column:is_builtin;default:0"`
	CreateAT      time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT      time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type ModelField struct {
	ID            uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	Name          string    `gorm:"type:VARCHAR(255) NOT NULL"`
	FieldUID      string    `gorm:"type:VARCHAR(255) NOT NULL;unique;uniqueIndex:uk_model_field_uid"`
	ModelUIDRefer string    `gorm:"type:VARCHAR(63) NOT NULL;column:model_uid"`
	IsEditable    bool      `gorm:"type:UNSIGNED TINYINT;default:1"`
	IsOptional    bool      `gorm:"type:UNSIGNED TINYINT;default:1"`
	Description   string    `gorm:"type:VARCHAR(1023)"`
	CreateAT      time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT      time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type ModelResource struct {
	ID                 uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	ModelUIDRefer      string    `gorm:"type:VARCHAR(63) NOT NULL;column:model_uid"`
	ResourceUID        string    `gorm:"type:VARCHAR(63) NOT NULL;uniqueIndex:uk_model_resource_uid"`
	CollectionUIDRefer string    `gorm:"type:VARCHAR(63);column:collection_uid"`
	IsAutoCollected    bool      `gorm:"type:UNSIGNED TINYINT;default:0"`
	CreateAT           time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT           time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type ModelResourceFieldData struct {
	ID               uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	Data             string    `gorm:"type:VARCHAR(1023) NOT NULL"`
	ModelUIDRefer    string    `gorm:"type:VARCHAR(63) NOT NULL;column:model_uid"`
	FieldUIDRefer    string    `gorm:"type:BIGINT UNSIGNED NOT NULL;column:field_uid;index:idx_model_resource_field_data"`
	ResourceUIDRefer string    `gorm:"type:BIGINT UNSIGNED NOT NULL;column:resource_uid;index:idx_model_resource_field_data"`
	CreateAT         time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT         time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type ModelDataCollection struct {
	ID             uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	Name           string    `gorm:"type:VARCHAR(255) NOT NULL"`
	ModelUIDRefer  string    `gorm:"type:VARCHAR(63) NOT NULL;column:model_uid"`
	CollectionUID  string    `gorm:"type:VARCHAR(63) NOT NULL;uniqueIndex:uk_model_data_collection_uid"`
	ScarpeEndpoint string    `gorm:"type:VARCHAR(63) NOT NULL"`
	CreateAT       time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT       time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type ModelAssociation struct {
	ID                        uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	AssociationUID            string    `gorm:"type:VARCHAR(63) NOT NULL;column:association_uid;uniqueIndex:uk_association_uid"`
	ModelUIDRefer             string    `gorm:"type:VARCHAR(63) NOT NULL;column:model_uid;index:idx_model_association"`
	AssociationModelUIDRefer  string    `gorm:"type:VARCHAR(63) NOT NULL;column:association_model_uid;index:idx_model_association"`
	AssociationStatusUIDRefer string    `gorm:"type:VARCHAR(63) NOT NULL;column:association_status_uid;index:idx_model_association"`
	CreateAT                  time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT                  time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}

type ModelAssociationStatus struct {
	ID                   uint      `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;primaryKey"`
	Name                 string    `gorm:"type:VARCHAR(255) NOT NULL"`
	AssociationStatusUID string    `gorm:"type:VARCHAR(63) NOT NULL;unique;uniqueIndex:uk_association_status_uid"`
	Description          string    `gorm:"type:VARCHAR(1023)"`
	CreateAT             time.Time `gorm:"type:TIMESTAMP;column:gmt_create;autoCreateTime:true"`
	UpdateAT             time.Time `gorm:"type:TIMESTAMP;column:gmt_modified;autoUpdateTime:true"`
}
