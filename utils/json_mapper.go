package utils

import (
	"math"
	"strconv"
	"time"
	"xiong/ball/model"
)

func Json2GameList(jsonRes map[string]interface{}) *model.GameList {
	results, ok := jsonRes["results"].([]interface{})
	if !ok {
		return nil
	}

	gameList := &model.GameList{}
	for _, sportInfo := range results {
		itemInfo := sportInfo.(map[string]interface{})
		id := itemInfo["id"].(string)
		league, ok := itemInfo["league"].(map[string]interface{})
		home, homeOK := itemInfo["home"].(map[string]interface{})
		away, awayOK := itemInfo["away"].(map[string]interface{})
		if ok && homeOK && awayOK {
			gameList.Info = append(
				gameList.Info,
				model.GameInfo{Name: league["name"].(string), GameId: id, HomeName: home["name"].(string), AwayName: away["name"].(string)},
			)
		}
	}
	return gameList
}

func Json2Game(gameId string, jsonRes map[string]interface{}) *model.Game {
	results := jsonRes["results"].([]interface{})

	items := results[0].([]interface{})
	var score float64
	hasScore := false
	var Quarter int32
	for _, result := range items {
		item := result.(map[string]interface{})

		if item["type"] == "PA" && item["OR"] == "1" && item["SU"] == "0" && item["HA"] != nil {
			haVal := item["HA"].(string)
			curVal, err := strconv.ParseFloat(haVal, 64)
			if err == nil {
				hasScore = true
				score = math.Max(score, curVal)
			}
		}

		if item["type"] == "EV" && item["CP"] != nil {
			if item["CP"] == "Q1" {
				Quarter = 1
			} else if item["CP"] == "Q2" {
				Quarter = 2
			} else if item["CP"] == "Q3" {
				Quarter = 3
			} else if item["CP"] == "Q4" {
				Quarter = 4
			} else {
				Quarter = 5
			}
		}

	}

	stats, ok := jsonRes["stats"].(map[string]interface{})
	if !ok {
		return nil
	}
	updateAt, err := strconv.ParseInt(stats["update_at"].(string), 10, 64)
	if err != nil {
		return nil
	}

	if !hasScore {
		return nil
	}
	return &model.Game{GameId: gameId, Time: time.Unix(updateAt, 0), TotalScore: float32(score), Quarter: Quarter}
}
