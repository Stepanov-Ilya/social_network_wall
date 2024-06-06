package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"social_network_wall/graph/model"
	"time"
)

// CreatePost is the aresolver for the cretePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, commentsEnabled bool) (*model.Post, error) {
	post := &model.Post{
		PostID:          uuid.New().String(),
		Title:           title,
		Content:         content,
		CreatedAt:       time.Now(),
		CommentsEnabled: commentsEnabled,
		Comments:        make([]*model.Comment, 0),
	}
	r.posts = append(r.posts, post)
	return post, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentCommentID *string, content string) (*model.Comment, error) {
	post, err := r.getPostById(ctx, postID)
	if err != nil {
		return nil, err
	}

	comment := &model.Comment{
		CommentID:          uuid.New().String(),
		PostID:             postID,
		ParentCommentID:    parentCommentID,
		Content:            content,
		CreatedAt:          time.Now(),
		ChildrenCommentsID: make([]*model.Comment, 0),
	}

	r.comments[postID] = append(r.comments[postID], comment)
	post.Comments = append(post.Comments, comment)

	if subs, ok := r.commentSubscriptions[postID]; ok {
		for _, ch := range subs {
			ch <- comment
		}
	}

	return comment, nil
}

// DisableComments is the resolver for the disableComments field.
func (r *mutationResolver) DisableComments(ctx context.Context, postID string) (*model.Post, error) {
	post, err := r.getPostById(ctx, postID)
	if err != nil {
		return nil, err
	}
	post.CommentsEnabled = false

	return post, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	post, err := r.getPostById(ctx, id)
	return post, err
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	commentChan := make(chan *model.Comment, 1)
	go func() {
		<-ctx.Done()
		close(commentChan)
	}()
	r.commentSubscriptions[postID] = append(r.commentSubscriptions[postID], commentChan)

	return commentChan, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

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
