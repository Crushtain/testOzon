// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID            string  `json:"id"`
	Author        string  `json:"author"`
	PublicationID string  `json:"publicationId"`
	ParentID      *string `json:"parentId,omitempty"`
	Content       string  `json:"content"`
}

type Mutation struct {
}

type Publication struct {
	ID               string     `json:"id"`
	Author           string     `json:"author"`
	Title            string     `json:"title"`
	Content          string     `json:"content"`
	Comments         []*Comment `json:"comments"`
	CommentsDisabled bool       `json:"commentsDisabled"`
}

type PublicationComments struct {
	Comments   []*Comment `json:"comments"`
	TotalCount int        `json:"totalCount"`
}

type Query struct {
}

type Subscription struct {
}