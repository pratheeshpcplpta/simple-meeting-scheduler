package routers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/database"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/helper"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/middlewares"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/models"
)

func LoginFormSubmit(c *gin.Context) {
	var form models.Login
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//
	// Check the
	//
	database.InitConnection()
	user, err := database.GetUserByUsername(form.Username)

	if err != nil || user.Password == "" {
		c.JSON(http.StatusOK, models.Response{
			Status:  "error",
			Message: "unable to find user details",
		})
		return
	}

	// Validate the user's password
	err = helper.ValidatePassword(form.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Status:  "error",
			Message: "unable to authenticate user. Check username and password",
		})
		return
	}

	// user login success
	encruserId := helper.Encrypt([]byte(user.Username), "auth_token")
	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "authentication success",
		Data: map[string][]byte{
			"token": encruserId,
		},
	})
}

//
// Schedule meeting
//
func ScheduleNewMeetings(c *gin.Context) {
	var form models.MeetingSchedulesRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeStampStartTime, _ := strconv.ParseUint("1", 10, 32)
	timeStampEndTime, _ := strconv.ParseUint("1", 10, 32)
	// var validation_error string

	db := database.InitConnection()
	err := db.Model(&models.MeetingSchedules{}).Create(&models.MeetingSchedules{
		Title:       form.Title,
		Description: form.Description,
		StartTime:   uint32(timeStampStartTime),
		EndTime:     uint32(timeStampEndTime),
		HostedBy:    1,
	})
	fmt.Println(err)
	c.JSON(http.StatusOK, "success")
	//
	//
	//
}

//
// UpcomingMeetings
//

func UpcomingMeetings(c *gin.Context) {
	result := map[string]interface{}{}

	db := database.InitConnection()
	db.Raw("SELECT * FROM meeting_schedules JOIN users ON meeting_schedules.hosted_by = users.id").Find(&result)

	fmt.Println(result)
	c.JSON(http.StatusOK, "success")
}

func Routes(route *gin.Engine) {
	v1 := route.Group("/api")
	{
		v1.Use(middlewares.APIMiddleware())
		v1.POST("/login", LoginFormSubmit)
		v1.POST("/schedule-meeting", ScheduleNewMeetings)
		v1.POST("/list-meetings/upcoming", UpcomingMeetings)
	}
}
