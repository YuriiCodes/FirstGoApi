package dbAPI

import (
	. "FirstAPI/models"
	"gorm.io/gorm"
)

func GetAllMessagesFromDb(db *gorm.DB) ([]Msg, error) {
	var messages []Msg

	if err := db.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func SendMessageToDb(db *gorm.DB, msg Msg) error {
	if err := db.Create(&msg).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMessagesToUserFromDb(db *gorm.DB, userID int) ([]Msg, error) {
	var msgs []Msg
	if err := db.Where("sender_id <> ?", userID).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func DeleteMessageFromDb(db *gorm.DB, msgID int) (Msg, error) {
	var msg Msg
	if err := db.First(&msg, msgID).Error; err != nil {
		return msg, err
	}

	if err := db.Delete(&msg, msgID).Error; err != nil {
		return msg, err
	}
	return msg, nil
}

func UpdateMessageInDb(db *gorm.DB, messageID int, editedMessage Msg) (Msg, error) {
	var msg Msg

	if err := db.First(&msg, messageID).Error; err != nil {
		return msg, err
	}

	if err := db.Model(&msg).Update("message", editedMessage.Message).Error; err != nil {
		return msg, err
	}
	return msg, nil
}
