package storage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/Crushtain/testOzon/graph/model"
	"github.com/Crushtain/testOzon/internal/database"
)

type PostgresStorage struct {
	db  *database.DB
	log *zap.Logger
}

func NewPostgresStorage(db *database.DB) *PostgresStorage {
	return &PostgresStorage{
		db:  db,
		log: zap.NewNop(),
	}
}

func (p *PostgresStorage) CreatePublication(publication *model.Publication) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := p.db.DB.Exec(ctx, `
	INSERT INTO publications (id, author, title, content, comments_disabled) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		publication.ID, publication.Author, publication.Title, publication.Content, publication.CommentsDisabled)
	if err != nil {
		return err
		p.log.Info("Error to create a post", zap.Error(err))
	}
	return nil
}

func (p *PostgresStorage) CreateComment(comment *model.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := p.db.DB.Exec(ctx, `
 	INSERT INTO comments (id, publication_id, author, content, parent_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		comment.ID, comment.PublicationID, comment.Author, comment.Content, comment.ParentID)

	if err != nil {
		p.log.Info("Error to create a comment", zap.Error(err))
		return err
	}
	return nil
}

func (p *PostgresStorage) GetAllPublications() ([]*model.Publication, error) {
	publications := make([]*model.Publication, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := p.db.DB.Query(ctx, `SELECT id, author, title, content, comments_disabled FROM publications`)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.log.Info("Strings are empty, no posts found", zap.Error(err))
			return []*model.Publication{}, nil
		}
		p.log.Info("Error to get all posts", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var publication model.Publication
		if err = rows.Scan(&publication.ID, &publication.Author, &publication.Title,
			&publication.CommentsDisabled); err != nil {
			p.log.Info("Error to put values into model.publication", zap.Error(err))
			return nil, err
		}
		publications = append(publications, &publication)
	}
	return publications, err
}

func (p *PostgresStorage) GetPublicationByID(id string) (*model.Publication, error) {
	var comments []*model.Comment
	var publication model.Publication
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := p.db.DB.Query(ctx,
		`
		SELECT 
			p.id, p.title, p.content, p.author, p.comments_disabled,
			c.id, c.publication_id, c.author, c.content, c.parent_id
		FROM 
			posts p
		LEFT JOIN 
			comments c 
		ON 
			p.id = c.publication_id
		WHERE 
			p.id = $1;
	`, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &model.Publication{}, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment

		if err = rows.Scan(&publication.ID, &publication.Title, &publication.Content, &publication.Author, &publication.CommentsDisabled,
			&comment.ID, &comment.PublicationID, &comment.Author, &comment.Content, &comment.ParentID); err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	publication.Comments = comments

	return &publication, nil

}

func (p *PostgresStorage) CommentsByPublicationID(publicationID string, limit *int, offset *int) ([]*model.Comment, error) {
	var comments []*model.Comment
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := p.db.DB.Query(ctx,
		`SELECT id, publication_id, author, content, parent_id
				FROM comments
					WHERE publication_id = $1
						LIMIT $2
							OFFSET $3`, publicationID, limit, offset)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*model.Comment{}, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		if err = rows.Scan(&comment.ID, &comment.PublicationID, &comment.Author, &comment.Content, &comment.ParentID); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (p *PostgresStorage) IsDisableComment(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result bool

	err := p.db.DB.QueryRow(ctx,
		`SELECT comments_disabled
				FROM publications 
					WHERE id = $1`, id).Scan(&result)
	if err != nil {
		p.log.Info("Error to check if comment is enable")
		return false, err
	}
	return result, nil

}
