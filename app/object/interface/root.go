package ObjectInterface

import "github.com/gin-gonic/gin"

type IObjectDelivery interface {
	Upload(c *gin.Context)
	Download(c *gin.Context)
}