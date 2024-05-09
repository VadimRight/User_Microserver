package main

import (
	"log"

	"github.com/VadimRight/User_Microserver/internal/config"
)

func main() {
	cfg := config.EnvLoad()
	log.Println(cfg)
	
}
