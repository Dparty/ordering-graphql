package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(inject *gorm.DB) {
	db = inject
}
