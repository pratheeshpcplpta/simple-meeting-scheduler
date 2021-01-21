package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pratheeshpcplpta/simple-meeting-scheduler/database"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/helper"
	"github.com/pratheeshpcplpta/simple-meeting-scheduler/models"
)

func main() {
	// TestEncrypDecryp()
	// Test()
	Test2()
}
func Test2() {
	time := time.Now()
	fmt.Println(time.Zone())
}
func TestEncrypDecryp() {
	enc := helper.AES_Encrypt("test", "login")
	fmt.Println(enc)

	dec := helper.AES_Decrypt(enc, "login")
	fmt.Println(dec)

	// enc := helper.Base64Encode([]byte("test123456789"))
	// fmt.Println(string(enc))

	// dec, _ := helper.Base64Decode([]byte(enc))
	// fmt.Println(string(dec))

}

func Test() {
	result := []map[string]interface{}{}

	meeting__details := make(map[string]map[string]interface{})
	db := database.InitConnection()
	db.
		Select([]string{
			"meeting_schedules.*",
		}).
		Model(&models.MeetingSchedules{}).
		Joins("JOIN user_meetings on user_meetings.mid=meeting_schedules.id").
		Where(map[string]interface{}{
			"user_meetings.uid": 1,
		}).
		Where("meeting_schedules.start_time < ?", time.Now().Unix()).
		Find(&result)

	list__mids := []string{}
	for _, _item := range result {
		str__id := strconv.Itoa(int(_item["id"].(uint)))
		list__mids = append(list__mids, str__id)

		meeting__details[str__id] = make(map[string]interface{}, 0)
		meeting__details[str__id] = _item
	}

	mid__participants := []map[string]interface{}{}
	db.
		Select([]string{
			"users.id,users.username,user_meetings.mid",
		}).
		Model(&models.UserMeetings{}).
		Joins("JOIN users on user_meetings.uid=users.id").
		Where("user_meetings.mid IN (" + strings.Join(list__mids, ",") + ")").
		Find(&mid__participants)

	for _, _item := range mid__participants {
		key := strconv.Itoa(_item["mid"].(int))
		if meeting__details[key]["participants"] != nil {
			meeting__details[key]["participants"] = make(map[string]string, 0)
		}
		meeting__details[key]["participants"] = map[string]string{
			key: _item["username"].(string),
		}
	}

	fmt.Println(meeting__details)
}
