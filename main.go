package main

import (
	"github.com/JooHyeongLee/go-docker-boilerplate/config"
	"github.com/JooHyeongLee/go-docker-boilerplate/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	ch := make(chan bool, 1)
	go initEnv(ch)

	// 모든 설정이 완료되었을 때 메인 시작
	select {
	case <-ch:
		log.Info("메인 시작")
	}
}

func initEnv(ch chan bool) {
	// config init
	config.LoadConfigration()
	// logger init
	utils.InitLogger()

	ch <- true
}