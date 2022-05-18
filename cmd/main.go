package main

import (
	. "FirstAPI/models"
	"github.com/spf13/viper"

	. "FirstAPI/db"
	. "FirstAPI/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	dbUrl := viper.GetString("database.url")

	// connect to DB
	db := Init(dbUrl)
	err = db.AutoMigrate(&User{}, &Msg{})
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

	router.Run(serverHost + ":" + serverPort)
}
