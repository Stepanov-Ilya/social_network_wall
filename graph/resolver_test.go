package graph

import (
	"context"
	"testing"
	"time"

	"social_network_wall/graph/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockPosts = []*model.Post{
	{
		PostID:          "post-1",
		Title:           "Title 1",
		Content:         "Content 1",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	},
	{
		PostID:          "post-2",
		Title:           "Title 2",
		Content:         "Content 2",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	},
}

func TestCreatePost(t *testing.T) {
	r := &Resolver{
		posts:    []*model.Post{},
		comments: make(map[string][]*model.Comment),
	}

	title := "Test Title"
	content := "Test Content"
	commentsEnabled := true

	post, err := r.Mutation().CreatePost(context.Background(), title, content, commentsEnabled)
	require.NoError(t, err)

	assert.Equal(t, title, post.Title)
	assert.Equal(t, content, post.Content)
	assert.Equal(t, commentsEnabled, post.CommentsEnabled)
	assert.NotNil(t, post.PostID)
	assert.WithinDuration(t, time.Now(), post.CreatedAt, time.Second)
}

func TestCreateComment(t *testing.T) {
	r := &Resolver{
		posts:    []*model.Post{},
		comments: make(map[string][]*model.Comment),
	}

	post := &model.Post{
		PostID:          "post-1",
		Title:           "Test Title",
		Content:         "Test Content",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	}
	r.posts = append(r.posts, post)

	content := "Test Comment Content"
	parentCommentID := "parent-1"

	comment, err := r.Mutation().CreateComment(context.Background(), post.PostID, &parentCommentID, content)
	require.NoError(t, err)

	assert.Equal(t, post.PostID, comment.PostID)
	assert.Equal(t, parentCommentID, *comment.ParentCommentID)
	assert.Equal(t, content, comment.Content)
	assert.NotNil(t, comment.CommentID)
	assert.WithinDuration(t, time.Now(), comment.CreatedAt, time.Second)
}

func TestDisableComments(t *testing.T) {
	r := &Resolver{
		posts:    []*model.Post{},
		comments: make(map[string][]*model.Comment),
	}

	post := &model.Post{
		PostID:          "post-1",
		Title:           "Test Title",
		Content:         "Test Content",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	}
	r.posts = append(r.posts, post)

	updatedPost, err := r.Mutation().DisableComments(context.Background(), post.PostID)
	require.NoError(t, err)

	assert.Equal(t, post.PostID, updatedPost.PostID)
	assert.False(t, updatedPost.CommentsEnabled)
}

func TestPosts(t *testing.T) {
	r := &Resolver{
		posts:    []*model.Post{},
		comments: make(map[string][]*model.Comment),
	}

	post1 := &model.Post{
		PostID:          "post-1",
		Title:           "Test Title 1",
		Content:         "Test Content 1",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	}
	post2 := &model.Post{
		PostID:          "post-2",
		Title:           "Test Title 2",
		Content:         "Test Content 2",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	}
	r.posts = append(r.posts, post1, post2)

	posts, err := r.Query().Posts(context.Background())
	require.NoError(t, err)

	assert.Len(t, posts, 2)
	assert.Equal(t, post1.PostID, posts[0].PostID)
	assert.Equal(t, post2.PostID, posts[1].PostID)
}

func TestPost(t *testing.T) {
	r := &Resolver{
		posts:    []*model.Post{},
		comments: make(map[string][]*model.Comment),
	}

	post := &model.Post{
		PostID:          "post-1",
		Title:           "Test Title",
		Content:         "Test Content",
		CreatedAt:       time.Now(),
		CommentsEnabled: true,
		Comments:        []*model.Comment{},
	}
	r.posts = append(r.posts, post)

	resultPost, err := r.Query().Post(context.Background(), post.PostID)
	require.NoError(t, err)

	assert.Equal(t, post.PostID, resultPost.PostID)
	assert.Equal(t, post.Title, resultPost.Title)
	assert.Equal(t, post.Content, resultPost.Content)
}
