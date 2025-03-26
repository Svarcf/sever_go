package main

import (
	"github.com/Svarcf/sever_go/internal/common"
	"github.com/Svarcf/sever_go/internal/config"
	"github.com/Svarcf/sever_go/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	db := common.InitDB()

	server := server.NewServer(db, cfg.APP_PORT)
	server.Start()
}
