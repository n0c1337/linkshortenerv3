package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Discriminator string
	Url           string
}
