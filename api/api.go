package api

import (
	"football/api/handlers"
	"football/config"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"     // swagger embed files
	// ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @description This is a sample article demo.
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	// docs.SwaggerInfo.Title = cfg.App
	// docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.POST("/create", h.CreateStation)
	r.GET("/getall", h.GetAll)
	// r.GET("/articlelist", h.GetById)

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
