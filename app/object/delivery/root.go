package ObjectDelivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	ImageInterface "ocr/app/image/interface"
	ImageRepository "ocr/app/image/repository"
	"ocr/model"
	"os"
	"strings"
	"time"
)

type ObjectDelivery struct {
	imageRepository ImageInterface.IImageRepository
}

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
	_uuid := uuid.New()
	dst := "./images/" + _uuid.String() + ".jpg"
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.String(500, fmt.Sprintf("write file failed"))
		return
	}

	// update database
	image := model.Image{
		UserId:    c.Request.Header.Get("user_id"), // get from jwt
		RequestId: c.Request.Header.Get("request_id"),
		Path:      dst,
		Status:    "pending",
		CreateAt:  time.Now().Format(time.RFC3339),
	}
	fmt.Println(q.imageRepository)
	_, err = q.imageRepository.InsertOne(image)
	if err != nil {
		c.String(500, fmt.Sprintf("db failed"))
		return
	}
	// send message
	// return result
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (q *ObjectDelivery) Download(c *gin.Context) {
	//imageId := c.Query("id")
	f, err := os.Open("images/aa22420b-cca9-4192-8c1d-1bf3a9d22a74.jpg")
	if err != nil {
		c.String(500, fmt.Sprintf("read file failed"))
	}
	defer f.Close()
	//bufio.NewReader(f)
	io.Copy(c.Writer, f)
}

func NewObjectDelivery() (*ObjectDelivery, error) {
	var q ObjectDelivery
	var err error
	q.imageRepository, err = ImageRepository.NewImageRepository()
	fmt.Println(q.imageRepository, err)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
