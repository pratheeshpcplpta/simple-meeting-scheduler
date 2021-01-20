package helper

import (
	"fmt"
	"testing"

	"github.com/pratheeshpcplpta/simple-meeting-scheduler/models"

	"github.com/pratheeshpcplpta/simple-meeting-scheduler/database"
)

func TestGeneratePassword(t *testing.T) {

	var str, hash string
	var err error

	fmt.Println("-------------------------------------------")
	str = "test"
	hash, err = GeneratePasswordHash(str)
	fmt.Println(str)
	fmt.Println(hash)
	fmt.Println(err)
	fmt.Println("-------------------------------------------")

	fmt.Println("-------------------------------------------")
	str = "user"
	hash, err = GeneratePasswordHash(str)
	fmt.Println(str)
	fmt.Println(hash)
	fmt.Println(err)
	fmt.Println("-------------------------------------------")

	fmt.Println("-------------------------------------------")
	str = "account"
	hash, err = GeneratePasswordHash(str)
	fmt.Println(str)
	fmt.Println(hash)
	fmt.Println(err)
	fmt.Println("-------------------------------------------")

}

func TestGenerateusers(t *testing.T) {
	GenerateusersFunc()
}

func GenerateusersFunc() {
	// var users = []models.Users{
	// 	{Username: "user"},
	// 	{Name: "jinzhu2"},
	// 	{Name: "jinzhu3"},
	// }

	users := map[int]map[string]string{
		0: map[string]string{
			"username": "username",
			"password": "username",
		},
		1: map[string]string{
			"username": "user1",
			"password": "user1",
		},
		2: map[string]string{
			"username": "user2",
			"password": "user2",
		},
		3: map[string]string{
			"username": "user3",
			"password": "user3",
		},
	}
	db := database.InitConnection()
	// var MUsers []models.Users
	for _, item := range users {
		hash, _ := GeneratePasswordHash(item["password"])
		// MUsers = append(MUsers, models.Users{Username: item["username"], Password: hash})
		db.Create(&models.Users{Username: item["username"], Password: hash})
	}
}
