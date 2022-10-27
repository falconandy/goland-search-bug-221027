package main

import (
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/falconandy/book/sqlite"
)

func NewBookCommand() *cli.Command {
	return &cli.Command{
		Name: "book",
		Subcommands: []*cli.Command{
			{
				Name: "add",
				Action: func(c *cli.Context) error {
					storage, err := sqlite.NewBookStorage("books.db")
					if err != nil {
						return err
					}
					defer func() { _ = storage.Close() }()

					err = storage.Migrate()
					if err != nil {
						return err
					}

					for _, arg := range c.Args().Slice() {
						bookTitle := filepath.Base(arg)
						ext := filepath.Ext(bookTitle)
						bookTitle = strings.TrimSuffix(bookTitle, ext)

						err = storage.SaveBook(c.Context, bookTitle, nil)
						if err != nil {
							return err
						}
					}

					return nil
				},
			},
			{
				Name: "delete",
				Action: func(c *cli.Context) error {
					storage, err := sqlite.NewBookStorage("books.db")
					if err != nil {
						return err
					}
					defer func() { _ = storage.Close() }()

					err = storage.Migrate()
					if err != nil {
						return err
					}

					for _, arg := range c.Args().Slice() {
						bookTitle := filepath.Base(arg)
						ext := filepath.Ext(bookTitle)
						bookTitle = strings.TrimSuffix(bookTitle, ext)

						err = storage.DeleteBook(c.Context, bookTitle)
						if err != nil {
							return err
						}
					}

					return nil
				},
			},
		},
	}
}
