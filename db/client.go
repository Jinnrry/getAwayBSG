package db

import (
	"context"
	"fmt"
	"getAwayBSG/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		client, _ := mongo.NewClient(options.Client().ApplyURI(configInfo["dburl"].(string) + "/" + configInfo["dbDatabase"].(string)))
		ctx := context.Background()
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

func GetLianjiaStatus() int {
	client := GetInstance().client
	ctx := GetInstance().ctx
	configInfo := configs.Config()
	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia_status := db.Collection("lianjia_status")
	var res bson.M
	err := lianjia_status.FindOne(ctx, bson.M{}).Decode(&res)
	if err != nil {
		return 0
	}

	index := res["index"].(int32)
	return int(index)
}

func SetLianjiaStatus(i int) {
	client := GetInstance().client
	ctx := GetInstance().ctx
	configInfo := configs.Config()
	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia_status := db.Collection("lianjia_status")
	lianjia_status.DeleteMany(ctx, bson.M{})
	lianjia_status.InsertOne(ctx, bson.M{"index": i})
}

func GetZhilianStatus() (int, int) {
	client := GetInstance().client
	ctx := GetInstance().ctx
	configInfo := configs.Config()
	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia_status := db.Collection("zhilian_status")
	var res bson.M

	var city_index int
	var kw_index int
	err := lianjia_status.FindOne(ctx, bson.M{}).Decode(&res)
	if err != nil {
		return 0, 0
	}
	if res["city_index"] != nil {
		city_index = int(res["city_index"].(int32))
	}

	if res["kw_index"] != nil {
		kw_index = int(res["kw_index"].(int32))
	}

	return city_index, kw_index
}

func SetZhilianStatus(cityIndex int, kwIndex int) {
	client := GetInstance().client
	ctx := GetInstance().ctx
	configInfo := configs.Config()
	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia_status := db.Collection("zhilian_status")
	lianjia_status.DeleteMany(ctx, bson.M{})
	lianjia_status.InsertOne(ctx, bson.M{"city_index": cityIndex, "kw_index": kwIndex})
}

func GetCtx() context.Context {
	return GetInstance().ctx
}

func GetClient() *mongo.Client {
	return GetInstance().client
}