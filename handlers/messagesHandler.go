package handlers

import (
	. "FirstAPI"
	. "FirstAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetAllMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		messagesFromBd, _ := GetAllMessagesFromDb(db)
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
		SendMessageToDb(db, newMsg)
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

		msgs, _ := GetAllMessagesToUserFromDb(db, id)

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

		msg, _ := DeleteMessageFromDb(db, id)

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

		msg, _ := UpdateMessageInDb(db, id, newMsg)

		c.IndentedJSON(http.StatusOK, msg)
		return

	}
}
