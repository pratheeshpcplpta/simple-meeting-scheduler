package models

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// meeting schedules request
type MeetingSchedulesRequest struct {
	Title        string
	Description  string
	StartTime    string
	EndTime      string
	Participants []string
}

type Response struct {
	Status  string
	Message string
	Data    interface{}
}

//
//
// Database models
//

// users
type Users struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Password     string
	ActiveStatus uint
}

// meeting schedules
type MeetingSchedules struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Description  string
	StartTime    uint32
	EndTime      uint32
	ActiveStatus uint
	HostedBy     uint
	User         Users `gorm:"foreignKey:HostedBy"`
}

// User meetings
type UserMeetings struct {
	ID uint `gorm:"primaryKey"`

	Uid  uint
	User Users `gorm:"foreignKey:Uid"`

	Mid      uint
	Meetings MeetingSchedules `gorm:"foreignKey:Mid"`
}
