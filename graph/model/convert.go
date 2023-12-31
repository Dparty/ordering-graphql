package model

import (
	"github.com/Dparty/common/utils"
	"github.com/Dparty/model/core"
	model "github.com/Dparty/model/restaurant"
	"github.com/chenyunda218/golambda"
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
	return nil
}

func UserBackward(user core.Account) User {
	return User{
		ID:    golambda.Reference(utils.UintToString(user.ID())),
		Email: &user.Email,
	}
}

func ItemInputConvert(input ItemInput) model.Item {
	return model.Item{
		Name: input.Name,
	}
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

func ItemBackward(item model.Item) Item {
	var i Item
	i.ID = utils.UintToString(item.ID)
	i.Name = item.Name
	i.Pricing = int(item.Pricing)
	return i
}
