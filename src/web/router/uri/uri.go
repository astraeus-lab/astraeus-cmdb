package uri

const (
	ModelRouteGroup = "/model"

	ModelListRoute             = "/list"
	ModelDetailListRoute       = "/detail/:modelUID"
	ModelMaintenanceByUIDRoute = "/maintenance/:modelUID"
)

const (
	ModelMaintenanceRouteGroup = "/maintenance"

	ModelFieldMaintenanceRoute       = "/filed/:modelUID"
	ModelResourceMaintenanceRoute    = "/resource/:modelUID"
	ModelAssociationMaintenanceRoute = "/association/:modelUID"
)

const (
	AssociationStatusRouteGroup = "/association"

	AssociationStatusListRoute             = "/list"
	AssociationStatusMaintenanceByUIDRoute = "/maintenance/:associationStatusUID"
)
