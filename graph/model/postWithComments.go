package model

type PostWithComments struct {
	Post     *Post      `json:"Post"`
	Comments []*Comment `json:"Comments"`
}
