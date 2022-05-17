package dbAPI

import (
	. "FirstAPI/models"
	"gorm.io/gorm"
)

func GetAllMessagesFromDb(db *gorm.DB) ([]Msg, error) {
	var messages []Msg
	db.Find(&messages)
	return messages, nil
}

func SendMessageToDb(db *gorm.DB, msg Msg) error {
	db.Create(&msg)
	return nil
}

func GetAllMessagesToUserFromDb(db *gorm.DB, userID int) ([]Msg, error) {
	var msgs []Msg
	db.Where("sender_id <> ?", userID).Find(&msgs)
	return msgs, nil
}

func DeleteMessageFromDb(db *gorm.DB, msgID int) (Msg, error) {
	var msg Msg
	db.First(&msg, msgID)
	db.Delete(&msg, msgID)
	return msg, nil
}

func UpdateMessageInDb(db *gorm.DB, messageID int, editedMessage Msg) (Msg, error) {
	var msg Msg
	db.First(&msg, messageID)
	db.Model(&msg).Update("message", editedMessage.Message)
	return msg, nil
}
