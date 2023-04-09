package models

import "time"

type User struct {
	ID                 int64  `gorm:"primaryKey" json:"id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Gender             string `json:"gender"`
	NumberDefaults     int32  `json:"number_defaults"`
	Permission         bool   `json:"permission"`
	AcceptNotification bool   `json:"accept_notification"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `sql:"index"`
}
