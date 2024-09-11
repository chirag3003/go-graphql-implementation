package main

import (
	"chirag3003/graphql-server/db"
	"chirag3003/graphql-server/graph/model"
	"context"
	"log"
	"net/http"
	"os"

	"chirag3003/graphql-server/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	database := &db.DB{
		Books: []*model.Book{},
	}

	srv.AddTransport(&transport.Websocket{}) // <---- This is the important part!
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middlewareOne(srv, database))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func middlewareOne(next http.Handler, db *db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("authorization")
		if header != "hello" {
			http.Error(w, "sahi id de", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "db", db)
		r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
