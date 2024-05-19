package api 

import (
	"github.com/gin-gonic/gin"
	"github.com/VadimRight/User_Microserver/internal/config"
	"fmt"
	"log"
)

func InitServer(cfg *config.Config) {
	const opt = "api.InitServer"
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.ServerPort))
	if err != nil {
		log.Fatalf("%s: %v", opt, err)
	}	
}
