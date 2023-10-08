package model

import (
	"github.com/Dparty/common/utils"
	model "github.com/Dparty/model/restaurant"
)

func Convert(from any) any {
	switch m := from.(type) {
	case model.Table:
		return TableBackward(m)
	case Table:
		return TableForward(m)
	case model.Restaurant:
		return RestaurantBackward(m)
	}
	return ""
}

func RestaurantBackward(restaurant model.Restaurant) Restaurant {
	return Restaurant{
		ID:           utils.UintToString(restaurant.ID),
		Name:         restaurant.Name,
		Description:  restaurant.Description,
		RestaurantId: restaurant.ID,
	}
}

func TableBackward(table model.Table) Table {
	return Table{
		Label:        table.Label,
		RestaurantId: table.RestaurantId,
	}
}

func TableForward(table Table) model.Table {
	t := model.Table{
		Label: table.Label,
	}
	t.ID = utils.StringToUint(table.ID)
	return t
}
