package models

import (
	"github.com/Nrdzz/iBooking/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

type Administrator struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Permissions bool   `json:"permissions"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:""index`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Administrator{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Seat{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Booking{})
}
