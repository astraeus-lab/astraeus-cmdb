package handler

import (
	"net/http"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model/v1"

	"github.com/gin-gonic/gin"
)

func GetModelFiledHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid, fieldUID := c.Param("modelUID"), c.Query("fieldUID")
	res, err := v1.NewModelField(uid, v1.NewOptions{}).RetrieveField(fieldUID)
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateModelFiledHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	source := &model.Field{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	uid := c.Param("modelUID")
	if err = v1.NewModelField(uid, v1.NewOptions{}).AddField(source); err != nil {
		if util.IsAlreadyExistErr(err) {
			c.JSON(http.StatusConflict, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func DeleteModelFiledHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid, fieldUID := c.Param("modelUID"), c.Query("fieldUID")
	if err = v1.NewModelField(uid, v1.NewOptions{}).DeleteField(fieldUID); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusGone, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func UpdateModelFiledHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	source := &model.Field{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	uid, fieldUID := c.Param("modelUID"), c.Query("fieldUID")
	if err = v1.NewModelField(uid, v1.NewOptions{}).UpdateFiled(fieldUID, source); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

/******************************************************************************/

func GetModelResourceHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid, resourceUID := c.Param("modelUID"), c.Query("resourceUID")
	res, err := v1.NewModelResource(uid, v1.NewOptions{}).RetrieveResource(resourceUID)
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateModelResourceHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	source := &model.Resource{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	uid := c.Param("modelUID")
	if err = v1.NewModelResource(uid, v1.NewOptions{}).AddResource(source); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func DeleteModelResourceHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid, resourceUID := c.Param("modelUID"), c.Query("resourceUID")
	if err = v1.NewModelResource(uid, v1.NewOptions{}).DeleteResource(resourceUID); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusGone, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func UpdateModelResourceHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	source := &model.Resource{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	uid, resourceUID := c.Param("modelUID"), c.Query("resourceUID")
	if err = v1.NewModelResource(uid, v1.NewOptions{}).UpdateResource(resourceUID, source); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

/******************************************************************************/

func GetModelAssociationHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid, associationStatusUID := c.Param("modelUID"), c.Query("associationStatusUID")
	res, err := v1.NewModelAssociation(uid, v1.NewOptions{}).RetrieveAssociation(associationStatusUID)
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateModelAssociationHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid := c.Param("modelUID")
	relateModelUID, associationStatusUID := c.Query("relateModelUID"), c.Query("associationStatusUID")
	err = v1.NewModelAssociation(uid, v1.NewOptions{}).AddAssociation(relateModelUID, associationStatusUID)
	if err != nil {
		if util.IsAlreadyExistErr(err) {
			c.JSON(http.StatusConflict, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func DeleteModelAssociationHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	uid, associationStatusUID := c.Param("modelUID"), c.Query("associationStatusUID")
	if err = v1.NewModelAssociation(uid, v1.NewOptions{}).DeleteAssociation(associationStatusUID); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusGone, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

//func UpdateModelAssociationHandler(c *gin.Context) {
//}
