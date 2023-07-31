package models

type Reply struct {
	ID            int    `gorm:"primaryKey" json:"-"`
	ResponderName string `json:"respondername"`
	ReplyContent  string `json:"replycontent"`
	Approved      bool   `json:"approved"`
	CommentId     int    `json:"commentid"`
	CreatedAt     int64  `json:"-"`
	UpdatedAt     int64  `json:"-"`
}
