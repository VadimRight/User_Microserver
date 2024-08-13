package router

import (
	"time"

	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
	"github.com/gin-gonic/gin"
)

func NewRegisterEndpoint(env config.EnvConfig, timeout time.Duration, db postgres.PostgresStorage, group gin.RouterGroup) {
}
