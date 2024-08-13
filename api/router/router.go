package router

import (
	"time"

	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
	"github.com/gin-gonic/gin"
)

func Setup(env config.Config, timeout time.Duration, db postgres.PostgresStorage, gin gin.Engine) {
	publicRouter := gin.Group("")
}
