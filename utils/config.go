package utils

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Tokens    []RateToken
	MongoHost string
}

func ParseConfig() AppConfig {
	bytes, err := os.ReadFile("./config.json")
	if err != nil {
		return AppConfig{}
	}
	var result map[string]interface{}
	if json.Unmarshal(bytes, &result) != nil {
		return AppConfig{}
	}

	appConfig := AppConfig{}
	tokens := result["tokens"].([]interface{})
	for _, tokenInfo := range tokens {
		tokenMap := tokenInfo.(map[string]interface{})
		token := tokenMap["token"].(string)
		limit := tokenMap["limitPerHour"].(float64)
		appConfig.Tokens = append(appConfig.Tokens, RateToken{Token: token, LimitPerHour: int(limit)})
	}
	mongoHost := result["mongoHost"].(string)
	appConfig.MongoHost = mongoHost
	return appConfig
}
