package main

import (
	. "FirstAPI"
	FirstAPI "FirstAPI/db"
	. "FirstAPI/handlers"
	"github.com/gin-gonic/gin"
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
