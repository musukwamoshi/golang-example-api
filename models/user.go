package models

//import guuid "github.com/google/uuid"

type User struct {
	ID        int    `gorm:"primaryKey" json:"-"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Hash      string `json:"hash"`
	IsAdmin   bool   `json:"isadmin"`
	CreatedAt int64  `json:"-"`
	UpdatedAt int64  `json:"-"`

	Sessions []Session `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	Articles []Article `gorm:"foreignKey:AuthorId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}
