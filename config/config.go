package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db struct {
		AcquireTimeout  int    `json:"acquireTimeout"`
		ConnectionLimit int    `json:"connectionLimit"`
		Timezone        string `json:"timezone"`
		Charset         string `json:"charset"`
		Database        string `json:"database"`
		Username        string `json:"username"`
		Password        string `json:"password"`
		Host            string `json:"host"`
		Port            string `json:"port"`
		Dialect         string `json:"dialect"`
		Define          struct {
		Timestamps  bool `json:"timestamps"`
		Underscored bool `json:"underscored"`
	} `json:"define"`
		Logging bool `json:"logging"`
	} `json:"db"`
		Crypto struct {
		Key string `json:"key"`
		Iv  string `json:"iv"`
	} `json:"crypto"`
		Path struct {
		Log string `json:"log"`
	} `json:"path"`
}

type Response struct {
	Result bool `json:"result"`
	Message string `json:"message"`
}

var config Config

func LoadConfigration() {
	var env = os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	file, err := os.Open("config/"+env+".json")

	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
}

func GetConfig() Config {
	return config
}
