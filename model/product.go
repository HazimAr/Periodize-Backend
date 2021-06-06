package model

import guuid "github.com/google/uuid"

type Product struct {
	ID        	int       	`gorm:"primaryKey" json:"id"`
	CreatedAt 	int64      	`gorm:"autoCreateTime" json:"-" `
	UpdatedAt 	int64      	`gorm:"autoUpdateTime" json:"-"`
	UserRefer 	guuid.UUID 	`json:"-"`
	Name      	string     	`json:"name"`
	Descripion	string     	`json:"description"`
	Private		bool		`gorm:"default:false" json:"private"`
	Experience 	[]string	`json:"experience"`
	Sport		[]string	`json:"sport"`
}
