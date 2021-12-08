package model

import "time"

type Game struct {
	Id         string    `bson:"_id,omitempty" json:"id"`
	GameId     string    `bson:"gameId" json:"gameId"`
	Time       time.Time `bson:"time" json:"time"`
	TotalScore float32   `bson:"totalScore" json:"totalScore"`
	Quarter    int32     `bson:"quarter" json:"quarter"`
}
