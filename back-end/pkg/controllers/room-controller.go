package controllers

import (
	"github.com/Nrdzz/iBooking/pkg/models"
	"github.com/Nrdzz/iBooking/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRoom(c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// log.Printf("%v", &json)

	room := &models.Room{
		ID:         utils.GetID(),
		RoomNumber: json["room_number"].(string),
		Location:   json["location"].(string),
		Total:      utils.Stoi(json["total"].(string), 16).(int16),
		Free:       utils.Stoi(json["total"].(string), 16).(int16),
	}

	if err := room.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": *room})
}

func GetRoom(c *gin.Context) {
	rooms, err := models.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}

func GetRoomById(c *gin.Context) {
	roomId := utils.Stoi(c.Param("roomId"), 64).(int64)
	room, _, err := models.GetRoomById(roomId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"room": room,
	})
}

func DeleteRoom(c *gin.Context) {
	roomId := utils.Stoi(c.Param("roomId"), 64).(int64)
	if err := models.DeleteSeatByRoomId(roomId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := models.DeleteRoom(roomId); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "deleteOK",
	})
}

// UpdateRoom delete and create new room for update
func UpdateRoom(c *gin.Context) {
	roomId := utils.Stoi(c.Param("roomId"), 64).(int64)
	if _, db, err := GetRoomById(roomId); err != nil {

	}
}
