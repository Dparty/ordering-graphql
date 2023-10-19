package model

import (
	"github.com/Dparty/common/utils"
	restaurantModel "github.com/Dparty/model/restaurant"
)

type Restaurant struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	RestaurantId uint   `json:"restaurantId"`
}

func (r *Restaurant) Entity() restaurantModel.Restaurant {
	var restaurant restaurantModel.Restaurant
	db.Find(&restaurant, utils.StringToUint(r.ID))
	return restaurant
}

func (r *Restaurant) Tables() []*Table {
	var ts []restaurantModel.Table
	db.Find(&ts, "restaurant_id = ?", r.ID)
	var tables []*Table
	for _, t := range ts {
		table := Convert(t).(Table)
		tables = append(tables, &table)
	}
	return tables
}

func (r *Restaurant) Items() []Item {
	return []Item{}
}

func (r *Restaurant) Owner() User {
	return UserBackward(r.Entity().Owner())
}
