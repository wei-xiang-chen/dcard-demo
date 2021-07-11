package main

import (
	"dcard/client"
	"dcard/controller/url"
	"dcard/middleware"
	"dcard/setting"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func init() {
	var err error

	err = setupSetting()
	if err != nil {
		log.Fatalf("inti.setupSetting err : %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("inti.setupDBSetting err : %v", err)
	}
}

func main() {

	router := initializeRoutes()
	http.ListenAndServe(":9000", router)
}

func initializeRoutes() http.Handler {

	router := gin.Default()
	router.Static("/api-docs", "./swagger/dist")

	router.GET("/:urlId", middleware.ErrorHandler(url.GetOriginal))

	v1Router := router.Group("/api/v1/")
	{
		urlsRouter := v1Router.Group("/urls/")
		{
			urlsRouter.POST("/", middleware.ErrorHandler(url.Transform))
		}
	}

	return cors.AllowAll().Handler(router)
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadRedisSetting()
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {

	client.RedisEngine = client.NewRedisEngine(setting.RsSetting)

	return nil
}
