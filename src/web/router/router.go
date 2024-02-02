package router

import (
	"github.com/astraeus-lab/astraeus-cmdb/src/web/router/handler"
	"github.com/astraeus-lab/astraeus-cmdb/src/web/router/uri"

	"github.com/gin-gonic/gin"
)

func RegistAllRoute(e *gin.Engine) {
	modelRouteGroup(e)
	modelMaintenanceRouteGroup(e)
	associationStatusRouteGroup(e)
}

func modelRouteGroup(e *gin.Engine) {
	group := e.Group(uri.ModelRouteGroup)
	{
		group.GET(uri.ModelListRoute, handler.ListModelHandler)
		group.GET(uri.ModelDetailListRoute, handler.ModelDetailListHandler)

		group.GET(uri.ModelMaintenanceByUIDRoute, handler.GetModelByUIDHandler)
		group.POST(uri.ModelMaintenanceByUIDRoute, handler.CreateModelHandler)
		group.DELETE(uri.ModelMaintenanceByUIDRoute, handler.DeleteModelByUIDHandler)
		group.PATCH(uri.ModelMaintenanceByUIDRoute, handler.UpdateModelByUIDHandler)
	}
}

func modelMaintenanceRouteGroup(e *gin.Engine) {
	group := e.Group(uri.ModelMaintenanceRouteGroup)
	{
		group.GET(uri.ModelFieldMaintenanceRoute, handler.GetModelFiledHandler)
		group.POST(uri.ModelFieldMaintenanceRoute, handler.CreateModelFiledHandler)
		group.DELETE(uri.ModelFieldMaintenanceRoute, handler.DeleteModelFiledHandler)
		group.PATCH(uri.ModelFieldMaintenanceRoute, handler.UpdateModelFiledHandler)

		group.GET(uri.ModelResourceMaintenanceRoute, handler.GetModelResourceHandler)
		group.POST(uri.ModelResourceMaintenanceRoute, handler.CreateModelResourceHandler)
		group.DELETE(uri.ModelResourceMaintenanceRoute, handler.DeleteModelResourceHandler)
		group.PATCH(uri.ModelResourceMaintenanceRoute, handler.UpdateModelResourceHandler)

		group.GET(uri.ModelAssociationMaintenanceRoute, handler.GetModelAssociationHandler)
		group.POST(uri.ModelAssociationMaintenanceRoute, handler.CreateModelAssociationHandler)
		group.DELETE(uri.ModelAssociationMaintenanceRoute, handler.DeleteModelAssociationHandler)
		//group.PATCH(uri.ModelAssociationMaintenanceRoute, handler.UpdateModelAssociationHandler)
	}
}

func associationStatusRouteGroup(e *gin.Engine) {
	group := e.Group(uri.AssociationStatusRouteGroup)
	{
		group.GET(uri.AssociationStatusListRoute, handler.ListAssociationStatusHandler)

		group.GET(uri.AssociationStatusMaintenanceByUIDRoute, handler.GetAssociationStatusHandler)
		group.POST(uri.AssociationStatusMaintenanceByUIDRoute, handler.CreateAssociationStatusHandler)
		group.DELETE(uri.AssociationStatusMaintenanceByUIDRoute, handler.DeleteAssociationStatusHandler)
		group.PATCH(uri.AssociationStatusMaintenanceByUIDRoute, handler.UpdateAssociationStatusHandler)
	}
}
