package ObjectDelivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageRepository "ocr.service.backend/app/image/repository"
	"ocr.service.backend/config"
	"ocr.service.backend/model"
	"ocr.service.backend/module"
	"os"
	"path"
	"strings"
	"time"
)

type ObjectDelivery struct {
	imageRepository ImageInterface.IImageRepository
}

var CONFIG, _ = config.NewConfig(nil)

func (q *ObjectDelivery) Upload(c *gin.Context) {
	// upload
	file, err := c.FormFile("file")
	// validate
	if err != nil {
		c.String(400, fmt.Sprintf("multipart/form-data require field 'file'"))
		return
	}
	if !strings.Contains(file.Header.Get("Content-Type"), "image") {
		c.String(400, fmt.Sprintf("please upload image"))
		return
	}
	// Upload the file to specific dst.
	_uuid := uuid.New().String()
	_md5 := module.GetMD5Hash(_uuid)
	imagePath := path.Join(CONFIG.GetString("OBJECT_PATH"), _md5[:2], _md5[2:3], _uuid+".jpg")
	err = os.MkdirAll(path.Dir(imagePath), 0777)
	if err != nil {
		c.String(500, fmt.Sprintf("mkdir failed, ")+err.Error())
		return
	}
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		c.String(500, fmt.Sprintf("write file failed"))
		return
	}

	// update database
	image := model.Image{
		Id:        uuid.New().String(),
		UserId:    c.Request.Header.Get("user_id"), // get from jwt
		RequestId: c.Request.Header.Get("request_id"),
		Path:      imagePath,
		Status:    "pending",
		CreateAt:  time.Now().Format(time.RFC3339),
	}
	_, err = q.imageRepository.InsertOne(image)
	if err != nil {
		c.String(500, fmt.Sprintf("db failed"))
		return
	}
	// send message
	// return result
	data, _ := json.Marshal(image)
	c.Writer.Write(data)
}

func (q *ObjectDelivery) Download(c *gin.Context) {
	imageId := c.Query("id")
	arrImage, err := q.imageRepository.Get(model.Image{Id: imageId})
	if err != nil {
		c.String(500, fmt.Sprintf("read from db failed"))
	}
	if len(arrImage) == 0 {
		c.String(404, fmt.Sprintf("not found"))
		return
	}
	f, err := os.Open(arrImage[0].Path)
	if err != nil {
		c.String(500, fmt.Sprintf("read file failed"))
	}
	defer f.Close()
	io.Copy(c.Writer, f)
}

func NewObjectDelivery() (*ObjectDelivery, error) {
	var q ObjectDelivery
	var err error
	ok, err := module.Exists(CONFIG.GetString("OBJECT_PATH"))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		err := os.MkdirAll(CONFIG.GetString("OBJECT_PATH"), 0777)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	q.imageRepository, err = ImageRepository.NewImageRepository()
	if err != nil {
		return nil, err
	}
	return &q, nil
}
