package storage

import "github.com/Crushtain/testOzon/graph/model"

type Storage interface {
	CreatePublication(publication *model.Publication) error
	CreateComment(comment *model.Comment) error
	GetAllPublications() ([]*model.Publication, error)
	GetPublicationByID(id string) (*model.Publication, error)
	CommentsByPublicationID(publicationID string, limit *int, offset *int) ([]*model.Comment, error)
	IsDisableComment(id string) (bool, error)
}
