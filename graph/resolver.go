package graph

import (
	"sync"

	"github.com/Crushtain/testOzon/graph/model"
	"github.com/Crushtain/testOzon/internal/storage"
)

type Resolver struct {
	Storage  storage.Storage
	Comments map[string][]chan *model.Comment
	Mu       sync.Mutex
}

func NewResolver(storage storage.Storage) *Resolver {
	return &Resolver{
		Storage:  storage,
		Comments: make(map[string][]chan *model.Comment),
	}
}
