package api

import (
	"fmt"
	"log"

	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/gin-gonic/gin"
)

type authMiddlewareHandler interface {
	Handler() gin.HandlerFunc
}

func InitServer(cfg config.Config, auth authMiddlewareHandler) {
	const opt = "api.InitServer"
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	publicRouter := r.Group("")
	protectedRouter := r.Group("")
	protectedRouter.Use(auth.Handler())
	r.Use(gin.Recovery())
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.ServerPort))
	if err != nil {
		log.Fatalf("%s: %v", opt, err)
	}
}
