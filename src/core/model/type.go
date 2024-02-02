package model

// Model resource data model storage format.
//
// Model should support JSON/YAML formatting
// for field sharing and custom data collector.
type Model struct {
	Metadata Metadata `json:"metadata" yaml:"metadata"`

	// Field is all information of the collection model, not storing actual content.
	// Key is the UID corresponding to the Field.
	//
	// Updating fields depends on the revision of the model, ensuring the
	// accuracy of model configuration. And when updating the fields,
	// should be treated as the overall metadata update of the model.
	Field map[string]*Field `json:"field" yaml:"field"`

	// Resource is all resource allocation items of the Model.
	// Key is the UID corresponding to the Resource.
	//
	// Resource is just a collection of all the field contents of the model,
	// which is divided into automatic collection and manual maintenance.
	Resource map[string]*Resource `json:"resource" yaml:"resource"`

	// Association is associated with other Model.
	// Key is the UID corresponding to the Association.
	Association map[string]*Association `json:"association" yaml:"association"`
}

type Metadata struct {
	// FieldUID global unique identifier of the model.
	//
	// Cannot be updated.
	UID string `json:"uid" yaml:"uid"`

	Name           string `json:"name" yaml:"name"`
	GroupUID       string `json:"groupUID" yaml:"groupUID"`
	IsBuiltin      bool   `json:"isBuiltin" yaml:"isBuiltin"`
	CreateTime     int64  `json:"createTime" yaml:"createTime"`
	LastUpdateTime int64  `json:"lastUpdateTime" yaml:"lastUpdateTime"`

	// Revision starts from 0 and records the updated version of the model,
	// which must be carried in both requests and responses.
	//
	// Before executing the update operation, it is necessary to first
	// obtain an existing revision. If it exceeds the amount carried by
	// the request, it will prompt to overwrite. If user agree to overwrite,
	// update the revision to the latest value(ensuring that SQL can execute
	// based on conditional judgment).
	//
	// When submitting an operation, SQL needs to add a revision equal to
	// the value obtained before executing and self increasing the revision.
	Revision string `json:"revision" yaml:"revision"`
}

type Field struct {
	Name string `json:"name" yaml:"name"`

	// FieldUID global unique field identification.
	//
	// Cannot be updated.
	FieldUID string `json:"fieldUID" yaml:"fieldUID"`

	IsEditable  bool   `json:"isEditable" yaml:"isEditable"`
	IsOptional  bool   `json:"isOptional" yaml:"isOptional"`
	Description string `json:"description" yaml:"description"`
}

type Resource struct {
	// ResourceUID global identification of resources.
	//
	// Cannot be updated.
	ResourceUID string `json:"resourceUID" yaml:"resourceUID"`

	// ResourceUID global identification of collection.
	//
	// Cannot be updated.
	DataCollectionUID string `json:"dataCollectionUID" yaml:"dataCollectionUID"`

	IsAutoCollected bool `json:"isAutoCollected" yaml:"isAutoCollected"`

	// FieldConenct store all field content as string data type.
	//
	// Any call or parsing at any position requires prior data conversion,
	// especially for integers and floating-point numbers.
	FieldConenct map[string]string `json:"fieldConenct" yaml:"fieldConenct"`
}

type Association struct {
	Name                 string   `json:"name" yaml:"name"`
	Description          string   `json:"description" yaml:"description"`
	AssociationStatusUID string   `json:"associationStatusUID" yaml:"associationStatusUID"`
	AssociationModelUID  []string `json:"associationModelUID" yaml:"associationModelUID"`
}

type AssociationStatus struct {
	Name                 string `json:"name" yaml:"name"`
	Description          string `json:"description" yaml:"description"`
	AssociationStatusUID string `json:"associationStatusUID" yaml:"associationStatusUID"`
}
