package models

type Message struct {
	Id        string
	Sender    uint
	Reciever  uint
	Message   string
	Timestamp uint
	Delivered bool
}
