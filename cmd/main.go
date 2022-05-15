package main

import (
	. "FirstAPI"
	FirstAPI "FirstAPI/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {

	db := FirstAPI.Init()
	err := db.AutoMigrate(&User{}, &Msg{})
	if err != nil {
		return
	}
	var users []User
	db.Find(&users)

	fmt.Println(users)

	router := gin.Default()

	// get all messages
	router.GET("/messages", func(c *gin.Context) {
		var messagesFromBd []Msg
		db.Find(&messagesFromBd)
		c.IndentedJSON(http.StatusOK, messagesFromBd)
	})

	// send new message
	router.POST("/messages", func(c *gin.Context) {
		var newMsg Msg
		if err := c.BindJSON(&newMsg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newMsg)
		c.IndentedJSON(http.StatusOK, newMsg)
	})

	// get all messages to user with specified ID
	router.GET("/messages/:id", func(c *gin.Context) {
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
	})

	// delete message by ID
	router.DELETE("/messages/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var msg Msg
		db.First(&msg, id)
		db.Delete(&msg, id)
		c.IndentedJSON(http.StatusOK, msg)
	})

	// update message by ID
	router.PUT("/messages/:id", func(c *gin.Context) {
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

	})

	/*   API for users   */

	// get all users
	router.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.IndentedJSON(http.StatusOK, users)
	})

	// create new user
	router.POST("/users", func(c *gin.Context) {
		var newUsr User
		if err := c.BindJSON(&newUsr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&newUsr)
		c.IndentedJSON(http.StatusOK, newUsr)
	})

	// delete user by ID
	router.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var usr User
		db.First(&usr, id)
		db.Delete(&usr, id)
		c.IndentedJSON(http.StatusOK, usr)
	})

	// update user name
	router.PUT("/users/:id", func(c *gin.Context) {
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
	})
	router.Run("localhost:8000")
}
