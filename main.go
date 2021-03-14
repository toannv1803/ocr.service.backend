package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs" // docs is generated by Swag CLI, you have to import it.
	ImageDelivery "ocr.service.backend/app/image/delivery"
	ObjectDelivery "ocr.service.backend/app/object/delivery"
	"ocr.service.backend/config"
	"os"
)

var CONFIG, _ = config.NewConfig(nil)

// @title OCR BACKEND API
// @version 1.0
func main() {
	router := gin.Default()
	objectDelivery, err := ObjectDelivery.NewObjectDelivery()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imageDelivery, err := ImageDelivery.NewImageDelivery()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	router.PUT("/api/v1/object", objectDelivery.Upload)
	router.POST("/api/v1/object", objectDelivery.Upload)
	router.GET("/api/v1/object", objectDelivery.Download)

	router.GET("/api/v1/images", imageDelivery.Gets)
	router.POST("/api/v1/images", imageDelivery.Update)

	// swagger
	router.GET("/swagger/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	url := ginSwagger.URL("/swagger/swagger.json") // The url pointing to API definition
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":" + CONFIG.GetString("NO_SSL_PORT"))
}
