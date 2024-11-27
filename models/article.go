package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title         string `binding:"required" json:"title" gorm:"not null;varchar(255)"`
	TitleSelector string `binding:"required" json:"titleSelector" gorm:"not null;"`
	Url           string `binding:"required" json:"url" gorm:"not null;`
}
type ArticleDTO struct {
	URLs          []string `binding:"required" json:"urls" `
	TitleSelector string   `binding:"required" json:"titleSelector"`
}
