package apicontent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apicontent "github.com/henrikroschmann/F-api-ker/domain/api-content"
	"github.com/henrikroschmann/F-api-ker/services"
	"github.com/henrikroschmann/F-api-ker/utils/errors"
)

func Save(c *gin.Context) {
	var jsons apicontent.Jsons

	if err := c.ShouldBindJSON(&jsons.Content); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	contentId, saveErr := services.CreateApiContent(jsons)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, contentId)
}

func Delete(c *gin.Context) {
	contentID := c.Param("content_id")

	saveErr := services.DeleteApiContent(contentID)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
}

func Get(c *gin.Context) {
	result, restErr := services.GetApiContent()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetById(c *gin.Context) {
	contentID := c.Param("content_id")
	result, restErr := services.GetApiContentById(contentID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result.Content)
}
