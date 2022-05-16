package handlers

import (
	. "FirstAPI"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetAllMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var messagesFromBd []Msg
		db.Find(&messagesFromBd)
		c.IndentedJSON(http.StatusOK, messagesFromBd)
	}
}

func SendMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newMsg Msg
		if err := c.BindJSON(&newMsg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newMsg)
		c.IndentedJSON(http.StatusOK, newMsg)
	}
}

func GetAllMessagesToUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var msgs []Msg
		db.Where("sender_id <> ?", id).Find(&msgs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, msgs)
	}
}

func DeleteMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var msg Msg
		db.First(&msg, id)
		db.Delete(&msg, id)
		c.IndentedJSON(http.StatusOK, msg)
	}
}

func UpdateMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		}
		var newMsg Msg
		if err := c.BindJSON(&newMsg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var msg Msg
		db.First(&msg, id)
		db.Model(&msg).Update("message", newMsg.Message)
		c.IndentedJSON(http.StatusOK, newMsg)
		return

	}
}
