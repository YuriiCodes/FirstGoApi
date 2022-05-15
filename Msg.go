package FirstAPI

type Msg struct {
	Id         int    `json:"id"  gorm:"primaryKey`
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
	Message    string `json:"message"`
}
