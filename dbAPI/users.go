package dbAPI

import (
	. "FirstAPI/models"
	"gorm.io/gorm"
)

func GetAllUsersFromDb(db *gorm.DB) ([]User, error) {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUserInDB(db *gorm.DB, newUser User) error {
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserFromDb(db *gorm.DB, userID int) (User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return user, err
	}
	if err := db.Delete(&user, userID).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUserInDb(db *gorm.DB, userID int, editedUser User) (User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return user, err
	}
	if err := db.Model(&user).Update("name", editedUser.Name).Error; err != nil {
		return user, err
	}
	return user, nil
}
