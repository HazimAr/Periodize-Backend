package model

import (
	guuid "github.com/google/uuid"
)

type User struct {
	Token     guuid.UUID `gorm:"primaryKey" json:"-"`
	ID	  	  guuid.UUID `gorm:"primaryKey" json:"-"`
	Image	  string 	 `json:"image"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Sessions  []Session  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	Products  []Product  `gorm:"foreignKey:UserRefer; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64      `gorm:"autoUpdateTime:milli" json:"-"`
}
