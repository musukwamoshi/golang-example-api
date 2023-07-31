package models

//import guuid "github.com/google/uuid"

type Comment struct {
	ID             int    `gorm:"primaryKey" json:"-"`
	CommenterName  string `json:"commentername"`
	CommentContent string `json:"commentcontent"`
	Approved       bool   ` json:"approved"`
	ArticleId      int    `json:"articleid"`
	CreatedAt      int64  `json:"-"`
	UpdatedAt      int64  `json:"-"`

	Replies []Reply `gorm:"foreignKey:CommentId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}
