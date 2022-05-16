package main

import (
	. "FirstAPI"
	FirstAPI "FirstAPI/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func main() {

	// connect to DB
	db := FirstAPI.Init()
	err := db.AutoMigrate(&User{}, &Msg{})
	if err != nil {
		return
	}

	router := gin.Default()

	/* API for messages */
	// get all messages
	router.GET("/messages", GetAllMessages(db))

	// send new message
	router.POST("/messages", SendMessage(db))

	// get all messages to user with specified ID
	router.GET("/messages/:id", GetAllMessagesToUser(db))

	// delete message by ID
	router.DELETE("/messages/:id", DeleteMessage(db))

	// update message by ID
	router.PUT("/messages/:id", UpdateMessage(db))

	/*   API for users   */
	// get all users
	router.GET("/users", GetAllUsers(db))

	// create new user
	router.POST("/users", CreateUser(db))

	// delete user by ID
	router.DELETE("/users/:id", DeleteUser(db))

	// update user name
	router.PUT("/users/:id", UpdateUser(db))

	router.Run("localhost:8000")
}

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

/* Hanlders for users: */
func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.IndentedJSON(http.StatusOK, users)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUsr User
		if err := c.BindJSON(&newUsr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newUsr)
		c.IndentedJSON(http.StatusOK, newUsr)
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var usr User
		db.First(&usr, id)
		db.Delete(&usr, id)
		c.IndentedJSON(http.StatusOK, usr)
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		}
		var newUsr User
		if err := c.BindJSON(&newUsr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userBeforeEditing User
		db.First(&userBeforeEditing, id)
		db.Model(&userBeforeEditing).Update("name", newUsr.Name)
		c.IndentedJSON(http.StatusOK, newUsr)
		return
	}
}
