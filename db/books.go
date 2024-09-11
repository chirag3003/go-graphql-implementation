package db

import "chirag3003/graphql-server/graph/model"

type DB struct {
	Books []*model.Book
}

func (d *DB) GetBooks() []*model.Book {
	return d.Books
}

func (d *DB) NewBook(book *model.Book) {
	d.Books = append(d.Books, book)
}
