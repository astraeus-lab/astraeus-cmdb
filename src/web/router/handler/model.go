package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model/v1"

	"github.com/gin-gonic/gin"
)

// TODO: how return message, json format?

func ListModelHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	res, err := v1.NewModelManager(v1.NewOptions{}).List()
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func ModelDetailListHandler(c *gin.Context) {
	var listErr error
	defer checkErr(c, listErr)

	detailType, modelUID := c.Query("type"), c.Param("modelUID")
	switch strings.ToUpper(detailType) {
	case "FIELD":
		listFiled, err := v1.NewModelField(modelUID, v1.NewOptions{}).ListFiled()
		if err != nil && !util.IsNotFoundErr(err) {
			listErr = err
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, listFiled)

	case "RESOURCE":
		listResource, err := v1.NewModelResource(modelUID, v1.NewOptions{}).ListResource()
		if err != nil && !util.IsNotFoundErr(err) {
			listErr = err
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, listResource)

	case "ASSOCIATION":
		listAssociation, err := v1.NewModelAssociation(modelUID, v1.NewOptions{}).ListAssociation()
		if err != nil && !util.IsNotFoundErr(err) {
			listErr = err
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, listAssociation)

	default:
		listErr = fmt.Errorf("listed type do not exist")
		c.JSON(http.StatusNotAcceptable, "")
	}
}

func GetModelByUIDHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	uid := c.Param("modelUID")
	withData := util.Str2Bool(c.Query("withData"))
	res, err := v1.NewModelManager(v1.NewOptions{}).Retrieve(uid, withData)
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateModelHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	source := &model.Metadata{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if err = v1.NewModelManager(v1.NewOptions{}).Create(source); err != nil {
		if util.IsAlreadyExistErr(err) {
			c.JSON(http.StatusConflict, "")
			return
		}
		if util.IsDataTypeErr(err) {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func DeleteModelByUIDHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	uid := c.Param("modelUID")
	if err = v1.NewModelManager(v1.NewOptions{}).Delete(uid); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusGone, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func UpdateModelByUIDHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	source := &model.Metadata{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	uid := c.Param("modelUID")
	if err = v1.NewModelManager(v1.NewOptions{}).Update(uid, source); err != nil {
		if util.IsNotFoundErr(err) || util.IsDataTypeErr(err) {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}
