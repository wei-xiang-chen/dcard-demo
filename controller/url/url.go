package url

import (
	"dcard/model"
	appError "dcard/model/error"
	"dcard/service/url_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Transform(c *gin.Context) error {

	var urlInput model.UrlInput

	c.ShouldBindJSON(&urlInput)

	if urlInput.Url == nil || urlInput.ExpireAt == nil {
		return appError.AppError{Message: "Check request body. Required fields are not filled."}
	}

	urlOutput, err := url_service.Transform(&urlInput)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, urlOutput)
	return nil
}

func GetOriginal(c *gin.Context) error {

	urlId := c.Param("urlId")

	originalUrl, err := url_service.GetOriginal(&urlId)
	if err != nil {
		return err
	}

	if originalUrl == nil {
		return appError.NotFoundError{}
	}

	c.Redirect(http.StatusMovedPermanently, *originalUrl)
	return nil
}
