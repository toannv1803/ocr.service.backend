package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs" // docs is generated by Swag CLI, you have to import it.
	ImageDelivery "ocr.service.backend/app/image/delivery"
	"ocr.service.backend/app/middleware"
	ObjectDelivery "ocr.service.backend/app/object/delivery"
	"ocr.service.backend/config"
	"os"
)

var CONFIG, _ = config.NewConfig(nil)

// @title OCR BACKEND API
// @version 1.0
func main() {
	router := gin.Default()
	router.Use(cors.Default())
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
	authMiddleware := middleware.NewAuth()
	routerApi := router.Group("/api/v1")
	{
		{ //No auth
			routerApi.GET("/object/:id", objectDelivery.DownloadById)
			routerApi.POST("/object", objectDelivery.Upload)
			routerApi.PUT("/object", objectDelivery.Upload)

			routerApi.GET("/image/:image_id", imageDelivery.GetById)
			routerApi.POST("/image/:image_id", imageDelivery.UpdateById)
			routerApi.GET("/images", imageDelivery.Gets)
		}

		routerApiAuth := routerApi.Group("/auth")
		routerApiAuth.Use(authMiddleware.MiddlewareFunc())
		{ // Auth
			routerApiAuth.GET("/object/:id", objectDelivery.DownloadById)
			routerApiAuth.POST("/object", objectDelivery.Upload)
			routerApiAuth.PUT("/object", objectDelivery.Upload)

			routerApiAuth.GET("/image/:image_id", imageDelivery.GetById)
			routerApiAuth.POST("/image/:image_id", imageDelivery.UpdateById)
			routerApiAuth.GET("/images", imageDelivery.Gets)
			routerApiAuth.DELETE("/images", imageDelivery.Delete)
			routerApiAuth.GET("/images/block-ids", imageDelivery.GetListBlockId)
		}
	}

	// swagger
	router.GET("/swagger/swagger.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	url := ginSwagger.URL("/swagger/swagger.json") // The url pointing to API definition
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	//	claims := jwt.ExtractClaims(c)
	//	log.Printf("NoRoute claims: %#v\n", claims)
	//	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	//})
	router.Run(":" + CONFIG.GetString("NO_SSL_PORT"))
}
