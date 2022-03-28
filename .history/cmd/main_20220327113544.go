package main

import (
	"gqlgen/cmd/graph"
	"gqlgen/cmd/graph/generated"
	"log"
	"net/http"
	"os"
	"gqlgen/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	sqlServer := &service.MySQLServer{
		Connected:  false,
		Dsn:        os.Getenv("MYSQL_DSN"),
		MaxRetries: 5,
		Conn:       nil,
	}

	services := &service.Services{
		Servers:   []service.ExternalServer{sqlServer},
		Failed:    []service.ExternalServer{},
		Connected: []service.ExternalServer{},
	}

	services.Connect()


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
