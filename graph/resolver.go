//go:generate go run github.com/99designs/gqlgen -v

package graph

import (
	"context"
	"fmt"
	"social_network_wall/graph/model"
)

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

func (r *Resolver) getPostById(ctx context.Context, postId string) (*model.Post, error) {
	for _, post := range r.posts {
		if post.PostID == postId {
			return post, nil
		}
	}

	return nil, fmt.Errorf("post not found")
}

func (r *Resolver) GetPostWithComments(ctx context.Context, postID string) (*model.PostWithComments, error) {
	post, err := r.getPostById(ctx, postID)
	if err != nil {
		return nil, err
	}

	comments := r.comments[postID]

	postWithComments := &model.PostWithComments{
		Post:     post,
		Comments: comments,
	}

	return postWithComments, nil
}
