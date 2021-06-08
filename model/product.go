package model

import guuid "github.com/google/uuid"

type Product struct {
	ID			guuid.UUID	`gorm:"primaryKey" json:"id"`
	CreatedAt 	int64      	`gorm:"autoCreateTime" json:"created"`
	UpdatedAt 	int64      	`gorm:"autoUpdateTime" json:"updated"`
	UserRefer 	guuid.UUID 	`json:"-"`
	Name      	string     	`json:"name"`
	Descripion	string     	`json:"description"`
	Private		bool		`json:"private"`
	Experience 	string		`json:"experience"`
	Sport		string		`json:"sport"`
}
