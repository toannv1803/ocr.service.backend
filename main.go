package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ImageDelivery "ocr.service.backend/app/image/delivery"
	ObjectDelivery "ocr.service.backend/app/object/delivery"
	"ocr.service.backend/config"
	"os"
)

var CONFIG, _ = config.NewConfig(nil)

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
	router.Run(":2020")
}
