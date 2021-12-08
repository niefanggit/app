package main

import (
	"xiong/ball/engine"
	"xiong/ball/utils"
)

func main() {
	appConfig := utils.ParseConfig()
	engine.EngineRun(appConfig.Tokens, appConfig.MongoHost)

	select {}
}
