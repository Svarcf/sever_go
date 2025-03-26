package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Svarcf/sever_go/internal/graph"
	"github.com/Svarcf/sever_go/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db   *gorm.DB
	port string
}

func NewServer(db *gorm.DB, port string) *Server {
	return &Server{db: db, port: port}
}

func login(c *gin.Context) {
	var loginCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginCredentials.Username == "admin" && loginCredentials.Password == "admin" {
		token, err := services.GenerateJWT(loginCredentials.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}

}

func (s *Server) Start() {
	userService := services.NewUserService(s.db)
	fileService := services.NewFileService(s.db)
	roleService := services.NewRoleService(s.db)
	toolTypeService := services.NewToolTypeService(s.db)
	mechanicalPressService := services.NewMechanicalPressService(s.db)
	rawMaterialService := services.NewRawMaterialService(s.db)
	standardPartService := services.NewStandardPartService(s.db)
	toolService := services.NewToolService(s.db)
	toolRepairRecordService := services.NewToolRepairRecordService(s.db)

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

	router := gin.Default()

	router.GET("/", func(c *gin.Context) { playground.Handler("GraphQL playground", "/query")(c.Writer, c.Request) })
	router.POST("/query", func(c *gin.Context) { srv.ServeHTTP(c.Writer, c.Request) })
	router.POST("/login", login)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.port)
	log.Fatal(router.Run(":" + s.port))

}
