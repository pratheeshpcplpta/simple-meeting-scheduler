package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/models"
)

var IgnoreMethods = []string{"GET", "HEAD", "OPTIONS"}

// validate api
type APIValidator struct {
	APIName string
	APIKey  string
}

// get the api params
func LoadAPIParams(c *gin.Context) APIValidator {
	r := c.Request

	apiVal := APIValidator{}

	if t := r.FormValue("_apikey"); len(t) > 0 {
		apiVal.APIKey = t
	} else if t := r.Header.Get("X-API-KEY"); len(t) > 0 {
		apiVal.APIKey = t
	}

	if t := r.FormValue("_apiname"); len(t) > 0 {
		apiVal.APIName = t
	} else if t := r.Header.Get("X-API-NAME"); len(t) > 0 {
		apiVal.APIName = t
	}

	return apiVal
}

// validate api params
func validate(params APIValidator) (status bool, err error) {

	if params.APIName == "" || params.APIKey == "" {
		err = fmt.Errorf("Unable to find api params")
		return
	}

	// validation with actual values comes here
	APIKEY := "AAhzclm9jHdypqdmEQx"
	APINAME := "api_thirdparty"

	if params.APIName != APINAME || params.APIKey != APIKEY {
		err = fmt.Errorf("Invalid api params")
		return
	}

	status = true
	return
}

// Middleware validates API token.
func APIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if InArray(IgnoreMethods, c.Request.Method) {
			c.Next()
			return
		}

		apiParams := LoadAPIParams(c)
		if ok, _ := validate(apiParams); !ok {
			c.JSON(400, models.Response{
				Status:  "error",
				Message: "API authentication failed",
				Data:    "",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func InArray(arr []string, value string) bool {
	inarr := false

	for _, v := range arr {
		if v == value {
			inarr = true
			break
		}
	}

	return inarr
}
