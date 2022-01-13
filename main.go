package main

import (
	"github.com/JooHyeongLee/go-docker-boilerplate/config"
	"github.com/JooHyeongLee/go-docker-boilerplate/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	// config init
	config.LoadConfigration()
	// logger init
	utils.InitLogger()
	log.Info("Main Func")
}
