package db

func (mg ModelGroup) TableName() string {

	return "model_group"
}

func (m Model) TableName() string {

	return "model"
}

func (mf ModelField) TableName() string {

	return "model_filed"
}

func (mr ModelResource) TableName() string {

	return "model_resource"
}

func (md ModelResourceFieldData) TableName() string {

	return "model_resource_filed_data"
}

func (mc ModelDataCollection) TableName() string {

	return "model_data_collection"
}

func (ma ModelAssociation) TableName() string {

	return "model_association"
}

func (mas ModelAssociationStatus) TableName() string {

	return "model_association_status"
}
