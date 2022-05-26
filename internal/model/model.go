package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Note struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
    Title string ` validate:"required,min=3,max=6"`
	SubTitle string
	Text string
}