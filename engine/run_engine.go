package engine

import (
	"context"
	"log"
	"time"
	"xiong/ball/service"
	"xiong/ball/utils"
)

type RunEngine struct {
	request     *utils.RateLimitRequest
	gameService *service.GameService
	cancelList  []context.CancelFunc
}

func (engine *RunEngine) startFetchGame(gameId string) {
	log.Printf("startFetchGame gameId %s ", gameId)
	result, err := engine.request.Request("https://api.b365api.com/v1/bet365/event", map[string]string{"FI": gameId})
	if err != nil {
		log.Printf("Request game: %v", err)
	}
	game := utils.Json2Game(gameId, result)
	if game != nil {
		engine.gameService.UpsertGame(game)
	}
}

func (engine *RunEngine) startInplayFilter() {
	log.Printf("Start fetch Inplay")
	result, err := engine.request.Request("https://api.b365api.com/v1/bet365/inplay_filter", map[string]string{"sport_id": "18"})
	if err != nil {
		log.Printf("Request filter: %v", err)
	}
	gameList := utils.Json2GameList(result)

	if gameList == nil {
		log.Fatalf("Fetch Game List Error: %v", result)
	}
	engine.gameService.UpsertGameList(gameList)

	for _, cancel := range engine.cancelList {
		cancel()
	}
	engine.cancelList = nil

	for _, gameInfo := range gameList.Info {
		gameId := gameInfo.GameId
		ctx, cancel := context.WithCancel(context.Background())
		go func(ctx context.Context, gameId string) {
			log.Printf("Start Fetch Game Info %s", gameId)
			ticker := time.NewTicker(time.Second)
			for {
				select {
				case <-ticker.C:
					engine.startFetchGame(gameId)
				case <-ctx.Done():
					log.Printf("Close Game Fetch %s", gameId)
					return
				}
			}
		}(ctx, gameId)
		engine.cancelList = append(engine.cancelList, cancel)
	}
	log.Printf("End fetch Inplay")
}

func (engine *RunEngine) Start(tokens []utils.RateToken, mongoHost string) {
	engine.request = utils.NewRateLimitRequest(tokens)

	engine.gameService = service.NewGameService(mongoHost)

	go func() {
		ticker := time.NewTicker(time.Minute * 2)
		for {
			engine.startInplayFilter()
			<-ticker.C
		}
	}()
}

func EngineRun(tokens []utils.RateToken, mongoHost string) {
	engine := &RunEngine{}
	engine.Start(tokens, mongoHost)
}
