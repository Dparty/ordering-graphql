package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/Dparty/common/utils"
	restaurantModel "github.com/Dparty/model/restaurant"
	"github.com/Dparty/ordering-graphql/graph/model"
)

// Owner is the resolver for the Owner field.
func (r *itemResolver) Owner(ctx context.Context, obj *model.Item) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - Owner"))
}

// CreateSession is the resolver for the createSession field.
func (r *mutationResolver) CreateSession(ctx context.Context, input *model.CreateSessionInput) (*model.Session, error) {
	panic(fmt.Errorf("not implemented: CreateSession - createSession"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	var item restaurantModel.Item
	c := db.Find(&item, utils.StringToUint(id))
	fmt.Println(item)
	if c.RowsAffected == 0 {
		return nil, nil
	}
	return &model.Item{
		Name: item.Name,
	}, nil
}

// Restaurant is the resolver for the restaurant field.
func (r *queryResolver) Restaurant(ctx context.Context, id string) (*model.Restaurant, error) {
	restaurant := restaurantModel.FindRestaurant(utils.StringToUint(id))
	if restaurant == nil {
		return nil, nil
	}
	return &model.Restaurant{
		ID:          utils.UintToString(restaurant.ID),
		Name:        restaurant.Name,
		Description: restaurant.Description,
	}, nil
}

// Table is the resolver for the table field.
func (r *queryResolver) Table(ctx context.Context, id string) (*model.Table, error) {
	panic(fmt.Errorf("not implemented: Table - table"))
}

// Item returns ItemResolver implementation.
func (r *Resolver) Item() ItemResolver { return &itemResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type itemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
