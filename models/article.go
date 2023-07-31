package models

type Article struct {
	ID       int    `gorm:"primaryKey" json:"-"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Approved bool   `json:"approved"`
	AuthorId int    `json:"authorid"`

	Comments []Comment `gorm:"foreignKey:ArticleId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}
