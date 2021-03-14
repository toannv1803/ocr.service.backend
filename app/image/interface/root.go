package ImageInterface

import (
	"github.com/gin-gonic/gin"
	"ocr/model"
)

type IImageDelivery interface {
	Gets(c *gin.Context)
	Update(c *gin.Context)
}

type IImageRepository interface {
	Get() (model.Image, error)
	InsertOne(image model.Image) (string, error)
	Update(filter model.Image, image model.Image) error
}
