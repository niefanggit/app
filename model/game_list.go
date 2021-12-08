package model

type GameInfo struct {
	Name     string `bson:"name" json:"name"`
	GameId   string `bson:"gameId" json:"gameId"`
	HomeName string `bson:"homeName" json:"homeName"`
	AwayName string `bson:"awayName" json:"awayName"`
}

type GameList struct {
	Id   string     `bson:"_id,omitempty" json:"id"`
	Info []GameInfo `bson:"info" json:"info"`
}
