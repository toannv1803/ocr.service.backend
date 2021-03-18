package ObjectInterface

import "github.com/gin-gonic/gin"

type IObjectDelivery interface {
	Upload(c *gin.Context)
	DownloadById(c *gin.Context)
}
