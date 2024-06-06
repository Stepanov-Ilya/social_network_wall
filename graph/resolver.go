//go:generate go run github.com/99designs/gqlgen -v

package graph

import "social_network_wall/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts                []*model.Post
	comments             map[string][]*model.Comment
	commentSubscriptions map[string][]chan *model.Comment
}

func newResolver() *Resolver {
	return &Resolver{
		posts:    []*model.Post{},
		comments: map[string][]*model.Comment{},
	}
}
