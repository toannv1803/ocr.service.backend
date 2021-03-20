package middleware

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"ocr.service.backend/config"
	"ocr.service.backend/enum"
	"ocr.service.backend/model"
	"time"
)

func NewAuth() *jwt.GinJWTMiddleware {
	CONFIG, _ := config.NewConfig(nil)
	var identityKey = CONFIG.GetString("IDENTITY_KEY")
	var secret = CONFIG.GetString("SECRET")
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secret),
		Timeout:     time.Hour,
		MaxRefresh:  2 * time.Hour,
		IdentityKey: identityKey,
		IdentityHandler: func(c *gin.Context) interface{} {
			fmt.Println("IdentityHandler")
			claims := jwt.ExtractClaims(c)
			return &model.Claim{
				UserId: claims[identityKey].(string),
				Role:   claims["role"].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			fmt.Println("Authorizator")
			if claim, ok := data.(*model.Claim); ok && (claim.Role == enum.RoleUser || claim.Role == enum.RoleAdmin) {
				agent := model.Agent{
					UserId: claim.UserId,
					Role:   claim.Role,
				}
				c.Set("agent", agent)
				fmt.Println(agent)
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Println("Unauthorized")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}
