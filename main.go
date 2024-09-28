package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Svarcf/sever_go/internal/common"
	"github.com/Svarcf/sever_go/internal/config"
	"github.com/Svarcf/sever_go/internal/graph"
	"github.com/Svarcf/sever_go/internal/services"
)

func main() {
	cfg := config.LoadConfig()
	db := common.InitDB()
	port := cfg.APP_PORT
	userService := services.NewUserService(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserService: userService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
