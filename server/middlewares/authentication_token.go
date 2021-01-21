package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/database"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/helper"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/models"
)

var authToken string

// get the api params
func LoadAccessToken(c *gin.Context) string {
	r := c.Request

	if t := r.FormValue("_access_token"); len(t) > 0 {
		authToken = t
	} else if t := r.Header.Get("X-ACCESS-TOKEN"); len(t) > 0 {
		authToken = t
	}
	return authToken
}

// Middleware validates API token.
func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if InArray(IgnoreMethods, c.Request.Method) {
			c.Next()
			return
		}

		access_token := LoadAccessToken(c)
		if access_token == "" {
			c.JSON(400, models.Response{
				Status:  "error",
				Message: "Access token authentication failed",
				Data:    "",
			})
			c.Abort()
			return
		}

		//
		// Validate token
		descrytID := helper.AES_Decrypt(access_token, "auth_token")
		database.InitConnection()
		user, err := database.GetUserByUsername(descrytID)
		if user.Username == "" || err != nil {
			c.JSON(400, models.Response{
				Status:  "error",
				Message: "Access token failure",
				Data:    "",
			})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
