package ImageInterface

import (
	"github.com/gin-gonic/gin"
	"ocr.service.backend/model"
)

type IImageDelivery interface {
	GetById(c *gin.Context)
	Gets(c *gin.Context)
	UpdateById(c *gin.Context)
	Delete(c *gin.Context)
	PublishTask(image model.Image) error
}

type IImageRepository interface {
	Get(filter model.Image) ([]model.Image, error)
	InsertOne(image model.Image) (string, error)
	Update(filter model.Image, image model.Image) (int64, error)
	Delete(filter model.Image) (int64, error)
}

type IImageUseCase interface {
	Gets(agent model.Agent, filter model.Image) ([]model.Image, error)
	InsertOne(agent model.Agent, image model.Image) (string, error)
	Update(agent model.Agent, filter model.Image, image model.Image) (int64, error)
	Delete(agent model.Agent, id model.Image) (int64, error)
}
