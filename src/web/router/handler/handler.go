package handler

import (
	"fmt"
	"strings"

	"github.com/astraeus-lab/astraeus-cmdb/src/core/model"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type postJSONData interface {
	*model.Model | *model.AssociationStatus |
		*model.Metadata | *model.Field | *model.Resource | *model.Association
}

func getPostJSONData[T postJSONData](c *gin.Context, out T) error {
	ct := c.Request.Header.Get("Content-Type")
	if strings.Compare(ct, strings.ToUpper("application/json")) != 0 {
		return fmt.Errorf("MIME type should be JSON, not %s", ct)
	}

	if err := c.ShouldBindJSON(out); err != nil {
		return err
	}

	return nil
}
