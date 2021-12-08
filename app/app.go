package main

import (
	"xiong/ball/service"
	"xiong/ball/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := utils.ParseConfig()
	service := service.NewGameService(appConfig.MongoHost)
	r := gin.Default()

	mwCORS := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
	})
	r.Use(mwCORS)

	r.GET("/gameList", func(c *gin.Context) {
		gameList := service.GetGameList()
		c.JSON(200, gameList)
	})

	r.GET("/game/:id", func(c *gin.Context) {
		gameId := c.Param("id")
		games := service.GetGameByGameId(gameId)
		c.JSON(200, games)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
