package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rezokk/graphql-sample/graph/generated"
	"github.com/rezokk/graphql-sample/graph/model"
)

var (
	satsu = "satsu"
	eco   = "eco"
	ry    = "ry"
	sh    = "sh"
	mi    = "mi"
	ritsu = "ritsu"
	jco   = "jco"
	aki   = "aki"
)

func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	people := []*model.Person{
		{
			Name:     ry,
			Parents:  []*string{&eco, &satsu},
			Children: []*string{&ritsu},
		},
		{
			Name:     mi,
			Parents:  []*string{&jco},
			Children: []*string{&ritsu},
		},
		{
			Name:     sh,
			Parents:  []*string{&eco, &satsu},
			Children: []*string{},
		},
		{
			Name:     satsu,
			Parents:  []*string{&aki},
			Children: []*string{&ry, &sh},
		},
		{
			Name:     eco,
			Parents:  []*string{},
			Children: []*string{&ry, &sh},
		},
		{
			Name:     ritsu,
			Parents:  []*string{&ry, &mi},
			Children: []*string{},
		},
		{
			Name:     jco,
			Parents:  []*string{},
			Children: []*string{&mi},
		},
		{
			Name:     aki,
			Parents:  []*string{},
			Children: []*string{&satsu},
		},
	}

	return people, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
