package handlers

import (
	. "FirstAPI"
	. "FirstAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

/* Hanlders for users: */
func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := GetAllUsersFromDb(db)
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUsr User
		if err := c.BindJSON(&newUsr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CreateUserInDB(db, newUsr)
		c.JSON(http.StatusOK, newUsr)
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		usr, _ := DeleteUserFromDb(db, id)
		c.JSON(http.StatusOK, usr)
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

		userAfterEditing, _ := UpdateUserInDb(db, id, newUsr)
		c.JSON(http.StatusOK, userAfterEditing)
		return
	}
}
