package route

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"github.com/Crushtain/testOzon/graph"
	"github.com/Crushtain/testOzon/internal/app"
)

const defaultPort = "8080"

func Route(app *app.App) http.Handler {
	router := chi.NewRouter()
	resolver := graph.NewResolver(app.Storage)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	return router
}
