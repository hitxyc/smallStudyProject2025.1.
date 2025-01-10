package main

import (
	"zhao_ci/api"
	"zhao_ci/config"
	"zhao_ci/mapper"
)

func main() {
	// @title 用户管理系统 API
	// @version 1.0
	// @description
	// @host localhost:8080
	// @BasePath
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	mapper.InitDataBase(&config.Database)
	defer mapper.CloseDataBase()
	r := api.InitRouter(mapper.GetDB())
	err = r.Run(config.Port)
	if err != nil {
		panic(err)
	}
}
