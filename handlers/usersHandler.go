package handlers

import (
	. "FirstAPI"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

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
