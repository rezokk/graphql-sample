package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/rezokk/graphql-sample/graph/generated"
	"github.com/rezokk/graphql-sample/graph/model"
)

func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	people := []string{
		"ry",
		"sh",
		"sat",
		"eco",
		"mi",
		"ritsu",
		"jko",
		"akio",
	}
	personModels := []*model.Person{}
	for i, p := range people {
		personModels = append(personModels, &model.Person{
			ID:   strconv.Itoa(i),
			Name: p,
		})
	}
	return personModels, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
