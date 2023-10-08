package graph

import (
	"github.com/Dparty/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(inject *gorm.DB) {
	db = inject
	model.Init(inject)
}
