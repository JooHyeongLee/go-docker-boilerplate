package utils

import (
	"fmt"
	"github.com/JooHyeongLee/go-docker-boilerplate/config"
	"io"
	"os"
	"time"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	// logger
	path := config.GetConfig().Path.Log
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", path, "%Y%m%d"),
		// rotatelogs.WithLinkName("./logs/current"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	log.SetOutput(io.MultiWriter(writer, os.Stdout))
}