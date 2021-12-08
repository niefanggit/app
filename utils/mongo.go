package utils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func NewDatabase(mongoHost string) *Database {
	var retryWrites bool = false

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().SetHosts([]string{mongoHost}).
		SetMaxPoolSize(8).
		SetRetryWrites(retryWrites)

	//设置用户名和密码
	// username := config.GetConf().MongoConf.Username
	// password := config.GetConf().MongoConf.Password

	// if len(username) > 0 && len(password) > 0 {
	//     clientOptions.SetAuth(options.Credential{Username: username, Password: password})
	// }

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &Database{Client: client}
}
