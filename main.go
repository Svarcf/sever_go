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
	fileService := services.NewFileService(db)
	roleService := services.NewRoleService(db)
	toolTypeService := services.NewToolTypeService(db)
	mechanicalPressService := services.NewMechanicalPressService(db)
	rawMaterialService := services.NewRawMaterialService(db)
	standardPartService := services.NewStandardPartService(db)
	toolService := services.NewToolService(db)
	toolRepairRecordService := services.NewToolRepairRecordService(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService:             userService,
		FileService:             fileService,
		RoleService:             roleService,
		ToolTypeService:         toolTypeService,
		MechanicalPressService:  mechanicalPressService,
		RawMaterialService:      rawMaterialService,
		StandardPartService:     standardPartService,
		ToolService:             toolService,
		ToolRepairRecordService: toolRepairRecordService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
