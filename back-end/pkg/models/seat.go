package models

import (
	"time"
)

type Seat struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	RoomID    int64 `json:"room_id"`
	X         int8  `json:"x"`
	Y         int8  `json:"y"`
	Status    int8  `json:"status"`
	Plug      bool  `json:"plug"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Seat) Create(RoomID int64) error {
	db.NewRecord(s)
	return db.Create(s).Error
}

func GetAllSeats() ([]Seat, error) {
	var seats []Seat
	if err := db.Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

func GetSeatById(id int64) (Seat, error) {
	var seats Seat
	if err := db.Where("ID =?", id).First(&seats).Error; err != nil {
		return Seat{}, err
	}
	return seats, nil
}

func DeleteSeatByRoomId(Id int64) error {
	var seat Seat
	return db.Model(&Seat{}).Where("room_id = ?", Id).Delete(&seat).Error
}

func DeleteSeat(id int64) error {
	var seat Seat
	return db.Where("ID =?", id).Delete(&seat).Error
}
