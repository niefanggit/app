package service

import (
	"context"
	"log"
	"xiong/ball/model"
	"xiong/ball/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameService struct {
	client             *utils.Database
	gameCollection     *mongo.Collection
	gameListCollection *mongo.Collection
}

func NewGameService(mongoHost string) *GameService {
	db := utils.NewDatabase(mongoHost)

	gameListCollection := db.Client.Database("game_bet").Collection("GameList")
	gameCollection := db.Client.Database("game_bet").Collection("Game")

	if _, err := gameCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{Keys: bson.D{{"gameId", -1}, {"time", -1}}}); err != nil {
		log.Fatalf("Create Index Error: %v", err)
	}
	return &GameService{client: db, gameCollection: gameCollection, gameListCollection: gameListCollection}
}

func (service *GameService) UpsertGameList(gameList *model.GameList) {
	if _, err := service.gameListCollection.DeleteOne(context.Background(), bson.M{}); err != nil {
		log.Fatalf("Delete game list error: %v", err)
	}

	service.gameListCollection.InsertOne(context.Background(), gameList)
}

func (service *GameService) UpsertGame(game *model.Game) {
	isUpsert := true
	_, err := service.gameCollection.ReplaceOne(context.Background(), bson.M{"gameId": game.GameId, "time": game.Time}, game, &options.ReplaceOptions{Upsert: &isUpsert})
	if err != nil {
		log.Fatalf("Upsert Game error: %v", err)
	}
}

func (service *GameService) GetGameList() *model.GameList {
	var gameList model.GameList
	err := service.gameListCollection.FindOne(context.Background(), bson.M{}).Decode(&gameList)
	if err != nil {
		log.Fatalf("Get GameList Error: %v", err)
	}
	return &gameList
}

func (service *GameService) GetGameByGameId(gameId string) []model.Game {
	cursor, err := service.gameCollection.Find(context.Background(), bson.M{"gameId": gameId}, &options.FindOptions{Sort: bson.D{{"time", 1}}})
	if err != nil {
		log.Fatalf("Get Games Error: %v", err)
	}
	defer cursor.Close(context.Background())

	var games []model.Game
	for cursor.Next(context.Background()) {
		var game model.Game
		if err = cursor.Decode(&game); err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}
	if len(games) == 0 {
		return games
	}
	var filterGames []model.Game
	for i, s := range games {
		if i == 0 || i == len(games)-1 {
			filterGames = append(filterGames, s)
		} else {
			var lastGameItem = filterGames[len(filterGames)-1]
			if lastGameItem.TotalScore != s.TotalScore || lastGameItem.Quarter != s.Quarter || s.Time.UnixMilli()-lastGameItem.Time.UnixMilli() >= 1000*59+500 {
				filterGames = append(filterGames, s)
			}
		}
	}
	return filterGames
}
