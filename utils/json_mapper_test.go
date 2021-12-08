package utils_test

import (
	"fmt"
	"testing"
	"xiong/ball/utils"
)

func TestJson2GameList(t *testing.T) {
	request := utils.NewRateLimitRequest([]utils.RateToken{
		{Token: "93784-6ES6B43suEWtTO", LimitPerHour: 1800},
	})

	result, err := request.Request("https://api.b365api.com/v1/bet365/inplay_filter", map[string]string{"sport_id": "18"})
	if err != nil {
		t.Fail()
	}

	gameList := utils.Json2GameList(result)
	if gameList == nil {
		t.Fail()
	}
	t.Logf("Get game list: %v", gameList)
}

func TestJson2Game(t *testing.T) {
	request := utils.NewRateLimitRequest([]utils.RateToken{
		{Token: "93784-6ES6B43suEWtTO", LimitPerHour: 1800},
	})

	result, err := request.Request("https://api.b365api.com/v1/bet365/event", map[string]string{"FI": "106665484"})
	if err != nil {
		t.Errorf("Request Error: %v", err)
	}

	game := utils.Json2Game("106696159", result)
	if game == nil {
		t.Errorf("Request Error: %v", game)
	}
	fmt.Printf("Game: %v", game)
	t.Fail()
}
