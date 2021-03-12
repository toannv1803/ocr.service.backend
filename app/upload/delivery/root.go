package upload_delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type UploadDelivery struct {
	r *gin.Engine
}

func (q UploadDelivery) UploadHandle(c *gin.Context) {
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
	// send message
	// return result
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func NewUploadDelivery() (*UploadDelivery, error) {
	return &UploadDelivery{}, nil
}
