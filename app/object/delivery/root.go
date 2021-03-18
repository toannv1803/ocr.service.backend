package ObjectDelivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	ImageDelivery "ocr.service.backend/app/image/delivery"
	ImageInterface "ocr.service.backend/app/image/interface"
	ImageUseCase "ocr.service.backend/app/image/usecase"
	ObjectInterface "ocr.service.backend/app/object/interface"
	"ocr.service.backend/config"
	"ocr.service.backend/model"
	"ocr.service.backend/module"
	"os"
	"path"
	"strings"
	"time"
)

type ObjectDelivery struct {
	imageDelivery ImageInterface.IImageDelivery
	imageUseCase  ImageInterface.IImageUseCase
}

var CONFIG, _ = config.NewConfig(nil)

// @tags Object
// @Summary upload, download object
// @Description upload object
// @start_time default
// @Param block_id header string false "add block id"
// @Param Authorization header string true "'Bearer ' + token"
// @Param file formData file true "add file multipart/form-data"
// @Success 200 {object} model.Image
// @failure 400 {string} string	"some info"
// @failure 404 {string} string	"not found"
// @failure 500 {string} string	"..."
// @Router /api/v1/auth/object [post]
func (q *ObjectDelivery) Upload(c *gin.Context) {
	// upload
	file, err := c.FormFile("file")
	// validate
	if err != nil {
		c.String(400, "multipart/form-data require field 'file'")
		return
	}
	if !strings.Contains(file.Header.Get("Content-Type"), "image") {
		c.String(400, "please upload image")
		return
	}
	// Upload the file to specific dst.
	_uuid := uuid.New().String()
	_md5 := module.GetMD5Hash(_uuid)
	imagePath := path.Join(CONFIG.GetString("OBJECT_PATH"), _md5[:2], _md5[2:3], _uuid+".jpg")
	err = os.MkdirAll(path.Dir(imagePath), 0777)
	if err != nil {
		c.String(500, "mkdir failed, "+err.Error())
		return
	}
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		c.String(500, "write file failed")
		return
	}
	userId := c.Request.Header.Get("user_id")
	if userId == "" {
		userId = "anonymous"
	}
	var agent model.Agent
	if v, ok := c.Get("agent"); ok {
		agent = v.(model.Agent)
		userId = agent.UserId
	}

	// update database
	image := model.Image{
		Id:       uuid.New().String(),
		UserId:   userId, // get from jwt
		BlockId:  c.Request.Header.Get("block_id"),
		Path:     imagePath,
		Status:   "pending",
		CreateAt: time.Now().Format(time.RFC3339),
	}
	_, err = q.imageUseCase.InsertOne(agent, image)
	if err != nil {
		c.String(500, "db failed")
		return
	}
	// send message
	err = q.imageDelivery.PublishTask(image)
	if err != nil {
		c.String(500, "rabbitmq failed")
		return
	}
	// return result
	c.JSON(200, image)
}

// @tags Object
// @Summary upload, download object
// @Description download object
// @start_time default
// @Param id path string true "object id"
// @Param Authorization header string true "'Bearer ' + token"
// @Success 200 {object} []byte
// @failure 400 {string} string	"some info"
// @failure 404 {string} string	"not found"
// @failure 500 {string} string	"..."
// @Router /api/v1/auth/object/{id} [get]
func (q *ObjectDelivery) DownloadById(c *gin.Context) {
	id := c.Param("id")
	if v, ok := c.Get("agent"); ok {
		agent := v.(model.Agent)
		arrImage, err := q.imageUseCase.Gets(agent, model.Image{Id: id})
		if err != nil {
			c.String(500, "read from db failed")
		}
		if len(arrImage) == 0 {
			c.String(404, "not found")
			return
		}
		f, err := os.Open(arrImage[0].Path)
		if err != nil {
			c.String(500, "read file failed")
		}
		defer f.Close()
		io.Copy(c.Writer, f)
	} else {
		c.String(400, "not allow")
	}

}
func NewObjectDelivery() (ObjectInterface.IObjectDelivery, error) {
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
	q.imageDelivery, err = ImageDelivery.NewImageDelivery()
	if err != nil {
		return nil, err
	}
	q.imageUseCase, err = ImageUseCase.NewImageUseCase()
	if err != nil {
		return nil, err
	}
	return &q, nil
}
