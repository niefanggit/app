package utils

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
)

type RateLimitRequest struct {
	rateLimiter []*rate.Limiter
	client      resty.Client
	tokenWaiter chan string
}

type RateToken struct {
	Token        string
	LimitPerHour int
}

func NewRateLimitRequest(tokens []RateToken) *RateLimitRequest {
	r := RateLimitRequest{}
	r.tokenWaiter = make(chan string)
	r.client = *resty.New()
	for _, t := range tokens {
		limiter := rate.NewLimiter(rate.Limit(t.LimitPerHour)/60/60, 2)
		r.rateLimiter = append(r.rateLimiter, limiter)
		go func(limiter *rate.Limiter, tokenWaiter chan string, rateToken RateToken) {
			for {
				c, _ := context.WithCancel(context.TODO())
				limiter.Wait(c)
				tokenWaiter <- rateToken.Token
			}
		}(limiter, r.tokenWaiter, t)
	}
	return &r
}

func (r *RateLimitRequest) Request(url string, query map[string]string) (map[string]interface{}, error) {
	token := <-r.tokenWaiter
	resp, err := r.client.R().SetQueryParam("token", token).SetQueryParams(query).Get(url)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &jsonRes); err != nil {
		return nil, err
	}

	return jsonRes, nil
}
