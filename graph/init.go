package graph

import (
	"github.com/Dparty/model"
	graphModel "github.com/Dparty/ordering-graphql/graph/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(inject *gorm.DB) {
	db = inject
	graphModel.Init(db)
	model.Init(inject)
}
