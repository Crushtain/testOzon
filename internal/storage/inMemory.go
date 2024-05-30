package storage

import (
	"errors"

	"github.com/Crushtain/testOzon/graph/model"
)

type InMemory struct {
	publications []*model.Publication
	comments     []*model.Comment
}

func NewInMemory() *InMemory {
	return &InMemory{
		publications: []*model.Publication{},
		comments:     []*model.Comment{},
	}
}
func (m *InMemory) CreatePublication(publication *model.Publication) error {
	m.publications = append(m.publications, publication)
	return nil
}
func (m *InMemory) CreateComment(comment *model.Comment) error {
	m.comments = append(m.comments, comment)
	return nil
}
func (m *InMemory) GetAllPublications() ([]*model.Publication, error) {
	if len(m.publications) == 0 {
		return []*model.Publication{}, errors.New("posts not found")
	}
	return m.publications, nil
}
func (m *InMemory) GetPublicationByID(id string) (*model.Publication, error) {
	var publication *model.Publication
	var comments []*model.Comment

	for _, p := range m.publications {
		if p.ID == id {
			publication = p
		}
	}
	if publication == nil {
		return &model.Publication{}, nil
	}

	for _, c := range m.comments {
		if c.PublicationID == id {
			comments = append(comments, c)
		}
	}
	publication.Comments = comments

	return publication, nil
}
func (m *InMemory) CommentsByPublicationID(publicationID string, limit *int, offset *int) ([]*model.Comment, error) {
	var comments []*model.Comment
	if *limit > len(m.comments) || *limit < 0 {
		return []*model.Comment{}, errors.New("limit exceeded")
	} else if *offset > len(m.comments) || *offset < 0 {
		return []*model.Comment{}, errors.New("offset exceeded")
	}

	for i := *offset; i < *limit; i++ {
		if m.comments[i].PublicationID == publicationID {
			comments = append(comments, m.comments[i])
		}
	}
	return comments, nil
}
func (m *InMemory) IsDisableComment(id string) (bool, error) {
	var post *model.Publication
	for _, currentPost := range m.publications {
		if currentPost.ID == id {
			post = currentPost
		}
	}

	if post == nil {
		return false, errors.New("post not found")
	}

	return post.CommentsDisabled, nil
}
