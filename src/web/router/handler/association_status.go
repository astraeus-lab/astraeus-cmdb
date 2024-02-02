package handler

import (
	"net/http"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"
	"github.com/astraeus-lab/astraeus-cmdb/src/core/model/v1"

	"github.com/gin-gonic/gin"
)

func ListAssociationStatusHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	res, err := v1.NewAssociationStatusManager(v1.NewOptions{}).List()
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetAssociationStatusHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	uid := c.Param("associationStatusUID")
	res, err := v1.NewAssociationStatusManager(v1.NewOptions{}).Retrivev(uid)
	if err != nil && !util.IsNotFoundErr(err) {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateAssociationStatusHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	source := &model.AssociationStatus{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if err = v1.NewAssociationStatusManager(v1.NewOptions{}).Create(source); err != nil {
		if util.IsAlreadyExistErr(err) {
			c.JSON(http.StatusConflict, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}
}

func DeleteAssociationStatusHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	uid := c.Param("associationStatusUID")
	if err = v1.NewAssociationStatusManager(v1.NewOptions{}).Delete(uid); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusGone, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}

func UpdateAssociationStatusHandler(c *gin.Context) {
	var err error
	defer checkErr(c, err)

	source := &model.AssociationStatus{}
	if err = getPostJSONData(c, source); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	uid := c.Param("associationStatusUID")
	if err = v1.NewAssociationStatusManager(v1.NewOptions{}).Update(uid, source); err != nil {
		if util.IsNotFoundErr(err) {
			c.JSON(http.StatusBadRequest, "")
			return
		}

		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusOK, "")
}
