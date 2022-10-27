package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"

	"github.com/falconandy/book"
)

func NewBookStorage(booksDBPath string) (anki.Storage, error) {
	db, err := sqlx.Connect("sqlite", booksDBPath)
	if err != nil {
		return nil, err
	}

	return &storage{
		db: db,
	}, nil
}

type storage struct {
	db *sqlx.DB
}

func (s *storage) Migrate() error {
	return nil
}

func (s *storage) Close() error {
	return nil
}

func (s *storage) SaveBook(ctx context.Context, bookTitle string, sentences []string) error {
	return nil
}

func (s *storage) DeleteBook(ctx context.Context, bookTitle string) error {
	return nil
}
