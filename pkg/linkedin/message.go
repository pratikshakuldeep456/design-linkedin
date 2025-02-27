package linkedin

type Message struct {
	ID         int
	Content    string
	ReceiverId int
	SenderId   int
}

func NewMessage(id, senderid, recid int, content string) *Message {
	return &Message{
		ID:         id,
		SenderId:   senderid,
		ReceiverId: recid,
		Content:    content,
	}
}
