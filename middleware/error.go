package middleware

import (
	api_err "dcard/model/error"
	"dcard/model/rest"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) error
type WsHandlerFunc func(c *gin.Context) (*int, error)

func ErrorHandler(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var restResult rest.RestResult
		var restError rest.RestError
		var err error

		err = handler(c)

		if err != nil {
			log.Printf("error: %v", err.Error())
			switch err.(type) {
			case api_err.AppError:
				restError.Message = err.Error()
				restResult.Error = &restError
				c.JSON(http.StatusBadRequest, restResult)
				return
			case api_err.NotFoundError:
				c.Status(http.StatusNotFound)
				return
			default:
				restError.Description = err.Error()
				restResult.Error = &restError
				c.JSON(http.StatusInternalServerError, restResult)
				return
			}
		}
	}
}
