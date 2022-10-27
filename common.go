package anki

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Migrate() error
	Close() error

	SaveBook(ctx context.Context, bookTitle string, sentences []string) error
	DeleteBook(ctx context.Context, bookTitle string) error
}

type Book struct {
	ID    uuid.UUID `db:"id"`
	Title string    `db:"title"`
}

type Sentence struct {
	BookID        uuid.UUID `db:"book_id"`
	SentenceIndex int       `db:"sentence_index"`
	Content       string    `db:"content"`
}
