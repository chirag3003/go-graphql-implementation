package graph

import (
	"chirag3003/graphql-server/db"
	"chirag3003/graphql-server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	database db.DB
	channels []chan []*model.Book
}

func (r *Resolver) NewChannel(ch chan []*model.Book) {
	r.channels = append(r.channels, ch)
}
