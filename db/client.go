package db

import (
	"context"
	"fmt"
	"getAwayBSG/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type singleton struct {
	client *mongo.Client
	ctx    context.Context
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
		configInfo := configs.Config()
		client, _ := mongo.NewClient(options.Client().ApplyURI(configInfo["dburl"].(string)))
		ctx, _ := context.WithTimeout(context.Background(), 3600*time.Second)
		instance.client = client
		instance.ctx = ctx
		err := client.Connect(ctx)
		if err != nil {
			fmt.Print("数据库连接失败！")
			fmt.Println(err)
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Print("ping error:")
			fmt.Println(err)
		}

	}
	return instance
}
