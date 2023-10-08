package model

import (
	model "github.com/Dparty/model/restaurant"
)

type Table struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

func (t *Table) Restaurant() Restaurant {
	restaurant := model.FindRestaurant(t.ID)
	return Convert(*restaurant).(Restaurant)
}
