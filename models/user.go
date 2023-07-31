package models

//import guuid "github.com/google/uuid"

type User struct {
	ID                      int    `gorm:"primaryKey" json:"-"`
	Email                   string `json:"email"`
	FirstName               string `json:"firstname"`
	LastName                string `json:"lastname"`
	Hash                    string `json:"hash"`
	Salt                    string `json:"salt"`
	IsAdmin                 bool   `json:"isadmin"`
	PasswordResetToken      string `json:"-"`
	PasswordResetExpiration string `json:"-"`
	CreatedAt               int64  `json:"-"`
	UpdatedAt               int64  `json:"-"`

	Sessions []Session `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	Articles []Article `gorm:"foreignKey:AuthorId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}
