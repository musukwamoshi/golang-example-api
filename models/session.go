package models

type Session struct {
	ID        int    `gorm:"primaryKey" json:"-"`
	UserId    int    `json:"userid"`
	UserEmail string `json:"useremail"`
	CreatedAt int64  `json:"-"`
	UpdatedAt int64  `json:"-"`
}
