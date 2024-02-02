package model

// CIModelManager manager configuration item model.
type CIModelManager interface {
	List() (map[string]*Model, error)

	Create(metadata *Metadata) error

	// Update replace the metadata of the model.
	//
	// UID related update operations should not be supported to avoid state confusion.
	// UID modification should be achieved through reconstruction (both direct modification
	// and reconstruction will result in a large number of DB operations).
	//
	// The version operation decision related to revision should be
	// completed by the routing time, so this method only performs execution.
	Update(modelUID string, metadata *Metadata) error

	// Retrieve query model related data from DB.
	//
	// Default Only obtain metadata of the model to ensure overall performance,
	// other relevant information is only obtained and used separately.
	Retrieve(modelUID string, withData bool) (*Model, error)

	Delete(modelUID string) error // TODO: need support soft delete?

	SerializeExport(modelUID, exportType string, withData bool) (string, error)
	SerializeImport(source string) (*Model, error)
}

// CIModelField is all information of the collection model,
// not storing actual content.
type CIModelField interface {
	ListFiled() (map[string]*Field, error)

	AddField(field *Field) error
	UpdateFiled(fieldUID string, field *Field) error
	RetrieveField(fieldUID string) (*Field, error)
	DeleteField(fieldUID string) error
}

// CIModelResource is all resource allocation items of the model.
type CIModelResource interface {
	ListResource() (map[string]*Resource, error)

	AddResource(resource *Resource) error
	UpdateResource(resourceUID string, resource *Resource) error
	RetrieveResource(resourceUID string) (*Resource, error)
	DeleteResource(resourceUID string) error
}

// CIModelAssociation is associated with other model.
type CIModelAssociation interface {
	ListAssociation() (map[string]*Association, error)

	AddAssociation(relateModelUID, associationStatusUID string) error
	RetrieveAssociation(associationStatusUID string) (*Association, error)
	DeleteAssociation(associationUID string) error
}

type CIAssociationStatusManager interface {
	List() (map[string]*AssociationStatus, error)

	Create(source *AssociationStatus) error
	Update(associationStatusUID string, source *AssociationStatus) error
	Retrivev(associationStatusUID string) (*AssociationStatus, error)
	Delete(associationStatusUID string) error
}
