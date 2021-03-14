package ImageDelivery

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageRepository "ocr.service.backend/app/image/repository"
	"ocr.service.backend/model"
)

type ImageDelivery struct {
	repository ImageInterface.IImageRepository
}

func (q *ImageDelivery) Gets(c *gin.Context) {
	var filter model.Image
	if c.BindQuery(&filter) == nil {
		arrImage, err := q.repository.Get(filter)
		if err != nil {
			c.String(500, "get data failed")
			return
		}
		byteImage, _ := json.Marshal(arrImage)
		c.Writer.Write(byteImage)
		return
	} else {
		fmt.Println(c.BindQuery(&filter))
		c.String(500, "...")
		return
	}
}

func (q *ImageDelivery) Update(c *gin.Context) {
	imageID := c.Query("id")
	var update model.Image
	err := c.BindJSON(&update)
	if err == nil {
		err := q.repository.Update(model.Image{Id: imageID}, model.Image{Data: update.Data})
		if err != nil {
			c.String(500, "update data failed")
			return
		}
		c.Writer.WriteHeader(200)
		return
	} else {
		c.String(500, "parser body failed, "+err.Error())
		return
	}
}

func NewImageDelivery() (*ImageDelivery, error) {
	var q ImageDelivery
	var err error
	q.repository, err = ImageRepository.NewImageRepository()
	if err != nil {
		return nil, err
	}
	return &q, nil
}
