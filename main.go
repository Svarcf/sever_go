package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Svarcf/sever_go/internal/common"
	"github.com/Svarcf/sever_go/internal/config"
	"github.com/Svarcf/sever_go/internal/graph"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	common.InitDB()
	port := cfg.APP_PORT

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
