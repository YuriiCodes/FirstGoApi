package dbAPI

import (
	. "FirstAPI/models"
	"gorm.io/gorm"
)

func GetAllUsersFromDb(db *gorm.DB) ([]User, error) {
	var users []User
	db.Find(&users)
	return users, nil
}

func CreateUserInDB(db *gorm.DB, newUser User) error {
	db.Create(&newUser)
	return nil
}

func DeleteUserFromDb(db *gorm.DB, userID int) (User, error) {
	var user User
	db.First(&user, userID)
	db.Delete(&user, userID)
	return user, nil
}

func UpdateUserInDb(db *gorm.DB, userID int, editedUser User) (User, error) {
	var user User
	db.First(&user, userID)
	db.Model(&user).Update("name", editedUser.Name)
	return user, nil
}
