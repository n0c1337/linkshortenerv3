package models

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Discriminator string
	Expire        time.Duration
}
