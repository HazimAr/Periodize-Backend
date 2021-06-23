package model

import guuid "github.com/google/uuid"

type Program struct {
	ID			guuid.UUID	`gorm:"primaryKey" json:"id"`
	CreatedAt 	int64      	`gorm:"autoCreateTime" json:"created"`
	UpdatedAt 	int64      	`gorm:"autoUpdateTime" json:"updated"`
	UserRefer 	guuid.UUID 	`json:"-"`
	Title      	string     	`json:"title"`
	Descripion	string     	`json:"description"`
	// Preset		bool		`json:"private"`
	Days		[]*Day		`gorm:"-" json:"days"`
}
