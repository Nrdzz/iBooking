package controllers

import (
	"bytes"
	"fmt"
	"github.com/Nrdzz/iBooking/pkg/models"
	"github.com/Nrdzz/iBooking/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSeat(c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// log.Printf("%v", &json)

	var errorMsg bytes.Buffer

	var Seats []models.Seat

	for _, v := range json {
		m := v.(map[string]interface{})
		// Check if the room exists
		roomId := utils.Stoi(m["room_id"].(string), 64).(int64)
		if _, _, err := models.GetRoomById(roomId); err != nil {
			errorMsg.WriteString(fmt.Sprintf("%s\v", err.Error()))
			continue
		}
		seat := &models.Seat{
			ID:     utils.GetID(),
			RoomID: roomId,
			X:      utils.Stoi(m["x"].(string), 8).(int8),
			Y:      utils.Stoi(m["y"].(string), 8).(int8),
			Status: utils.Stoi(m["status"].(string), 8).(int8),
			Plug:   m["plug"].(bool),
		}
		if err := seat.Create(roomId); err != nil {
			errorMsg.WriteString(fmt.Sprintf("%s\v", err.Error()))
			continue
		}
		Seats = append(Seats, *seat)
	}
	if errorMsg.Len() > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorMsg.String(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"seats": Seats,
	})
}
