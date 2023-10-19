package model

import "github.com/Dparty/model/restaurant"

type User struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	// Restaurants []*Restaurant `json:"restaurants,omitempty"`
}

func (u User) Restaurants() []Restaurant {
	var rs []restaurant.Restaurant
	db.Find(&rs, "account_id = ?", u.ID)
	var restaurants []Restaurant
	for _, r := range rs {
		restaurants = append(restaurants, RestaurantBackward(r))
	}
	return restaurants
}
