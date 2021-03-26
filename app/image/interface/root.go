package ImageInterface

import (
	"github.com/gin-gonic/gin"
	"ocr.service.backend/model"
	"ocr.service.backend/module/db"
)

type IImageDelivery interface {
	GetById(c *gin.Context)
	Gets(c *gin.Context)
	GetListBlockId(c *gin.Context)
	UpdateById(c *gin.Context)
	Delete(c *gin.Context)
	PublishTask(image model.ImageTask) error
}

type IImageRepository interface {
	Find(filter model.Image, option db.FindOption) ([]model.Image, int64, error)
	FindCustom(filter model.Image, res interface{}, option db.FindOption) (int64, error)
	InsertOne(image model.Image) (string, error)
	Update(filter model.Image, image model.Image) (int64, error)
	Delete(filter model.Image) (int64, error)
	Distinct(field string, filter interface{}) ([]interface{}, error)
}

type GetOption struct {
	Limit int64
	Skip  int64
}

type IImageUseCase interface {
	Gets(agent model.Agent, filter model.Image, option GetOption) ([]model.Image, int64, error)
	GetsCustom(agent model.Agent, filter model.Image, res interface{}, option GetOption) (int64, error)
	InsertOne(agent model.Agent, image model.Image) (string, error)
	Update(agent model.Agent, filter model.Image, image model.Image) (int64, error)
	Delete(agent model.Agent, id model.Image) (int64, error)
	GetListBlockId(agent model.Agent) ([]string, error)
}
