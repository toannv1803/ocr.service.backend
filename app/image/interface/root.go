package ImageInterface

import (
	"github.com/gin-gonic/gin"
	"ocr.service.backend/model"
)

type IImageDelivery interface {
	Gets(c *gin.Context)
	Update(c *gin.Context)
}

type IImageRepository interface {
	Get(filter model.Image) ([]model.Image, error)
	InsertOne(image model.Image) (string, error)
	Update(filter model.Image, image model.Image) error
}
