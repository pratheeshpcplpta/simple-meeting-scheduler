package routers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	if form.Username == "" || form.Password == "" {
		c.JSON(http.StatusOK, models.Response{
			Status:  "error",
			Message: "Unable to find params/params missing",
		})
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
	encruserId := helper.AES_Encrypt(user.Username, "auth_token")

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "authentication success",
		Data: map[string]string{
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

	//ge the user detail from context
	user, _ := c.Get("user")

	//
	//
	//
	participants := strings.Split(form.Participants, ",")
	fmt.Println("participants----", participants)
	//select the not availble participants
	//
	// example, if new meeting is going to schedule 1 to 2
	//
	// Conditions check
	// 	- start time < new start time and end time >= new start time
	// 						OR
	//  - start time > new end time and endtime >= new end time

	layout := "2006-01-02T15:04:05"
	// date_only_layout := "2006-01-02"
	start_time_unix, err := time.Parse(layout, form.StartTime)

	if err != nil {
		fmt.Println(err)
	}

	end_time_unix, err := time.Parse(layout, form.EndTime)

	if err != nil {
		fmt.Println(err)
	}

	// start_time__date_only := start_time_unix.Format("2006-01-02")
	// start_time__date_only_t, err := time.Parse(date_only_layout, start_time__date_only)
	// start_time__date_only_unix := start_time__date_only_t.Unix()

	// end_time__date_only := end_time_unix.Format("2006-01-02")
	// end_time__date_only_t, err := time.Parse(date_only_layout, end_time__date_only)
	// end_time__date_only_unix := end_time__date_only_t.Unix()

	db := database.InitConnection()

	result := []map[string]interface{}{}

	db.
		Raw(`
			SELECT uid,username FROM user_meetings 
			JOIN users on users.id=user_meetings.uid
			JOIN meeting_schedules on meeting_schedules.id=user_meetings.mid
			WHERE (
				(
					(start_time <= ? AND end_time >= ?)
							OR
					(start_time >= ? AND end_time <= ? )
				)
						AND
				(user_meetings.uid IN (?))
			)
		`, start_time_unix.Unix(), start_time_unix.Unix(), end_time_unix.Unix(), end_time_unix.Unix(), strings.Join(participants, ",")).
		Find(&result)

	if len(result) > 0 {
		var existing_meet_users map[int]string
		existing_meet_users = make(map[int]string, 0)
		for _, _item := range result {
			uid__str := int(_item["uid"].(int64))
			existing_meet_users[uid__str] = _item["username"].(string)
		}
		c.JSON(http.StatusOK, models.Response{
			Status:  "error",
			Message: "Scheduler failed. Some of the user already have meeting within the selected time slot",
			Data:    existing_meet_users,
		})
		return
	}

	Schedule := models.MeetingSchedules{
		MeetingId:   "MCSH-" + strings.ToUpper(helper.GenerateRandomString(6)),
		Title:       form.Title,
		Description: form.Description,
		StartTime:   uint32(start_time_unix.Unix()),
		EndTime:     uint32(end_time_unix.Unix()),
		HostedBy:    user.(models.Users).ID,
	}
	db.Create(&Schedule)

	//
	// Insert the participants to table with meeting id
	//
	var UserMeetings []models.UserMeetings
	UserMeetings = make([]models.UserMeetings, len(participants))

	var key int
	for _, participant := range participants {
		p__int, _ := strconv.Atoi(participant)
		UserMeetings[key] = models.UserMeetings{
			Mid: int(Schedule.ID),
			Uid: p__int,
		}
		key++
	}
	if len(UserMeetings) > 0 {
		db.Create(&UserMeetings)
	}

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Scheduler success",
		Data:    "",
	})
}

//
// UpcomingMeetings
//
func UpcomingMeetings(c *gin.Context) {
	result := []map[string]interface{}{}

	db := database.InitConnection()

	user, _ := c.Get("user")
	if user.(models.Users).ID > 0 {
		db.
			Select([]string{"meeting_schedules.*"}).
			Model(&models.MeetingSchedules{}).
			Joins("JOIN user_meetings on user_meetings.mid=meeting_schedules.id").
			Where(map[string]interface{}{
				"user_meetings.uid": user.(models.Users).ID,
			}).
			Where("meeting_schedules.start_time >= ?", time.Now().Unix()).
			Find(&result)

		c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "",
			Data:    result,
		})
	}

}

//
// UpcomingMeetings
//
func RecentMeetings(c *gin.Context) {
	result := []map[string]interface{}{}

	db := database.InitConnection()

	user, _ := c.Get("user")
	if user.(models.Users).ID > 0 {
		db.
			Select([]string{"meeting_schedules.*"}).
			Model(&models.MeetingSchedules{}).
			Joins("JOIN user_meetings on user_meetings.mid=meeting_schedules.id").
			Where(map[string]interface{}{
				"user_meetings.uid": user.(models.Users).ID,
			}).
			Where("meeting_schedules.start_time < ?", time.Now().Unix()).
			Find(&result)

		c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "",
			Data:    result,
		})
	}

}

//
// GetUsers
//

func GetUsers(c *gin.Context) {
	result := []map[string]interface{}{}
	db := database.InitConnection()
	db.
		Model(&models.Users{}).
		Select([]string{"username", "id"}).
		Find(&result)

	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "",
		Data:    result,
	})
}

func Routes(route *gin.Engine) {

	route.GET("/", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "login.tpl", gin.H{})
	})

	route.GET("/logout", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "login.tpl", gin.H{})
	})

	route.GET("/dashboard", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "dashboard.tpl", gin.H{})
	})

	v1 := route.Group("/api")
	{
		v1.Use(middlewares.APIMiddleware())
		v1.POST("/login", LoginFormSubmit)

		v1.Use(middlewares.AuthTokenMiddleware())
		v1.POST("/schedule-meeting", ScheduleNewMeetings)
		v1.POST("/list-meetings/upcoming", UpcomingMeetings)
		v1.POST("/list-meetings/recent", RecentMeetings)
		v1.POST("/get-users", GetUsers)
	}
}
