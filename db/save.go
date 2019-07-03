package db

import (
	"fmt"
	"getAwayBSG/configs"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func Add(item bson.M) {
	configInfo := configs.Config()
	client := GetInstance().client
	ctx := GetInstance().ctx

	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia := db.Collection(configInfo["dbCollection"].(string))
	_, err := lianjia.InsertOne(ctx, item)
	if err != nil {
		if !strings.Contains(err.Error(), "multiple write errors") {
			fmt.Print("数据库插入失败！")
			fmt.Println(err)
		}
	}

}

func Update(link string, m bson.M) {
	configInfo := configs.Config()

	client := GetInstance().client
	ctx := GetInstance().ctx

	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia := db.Collection(configInfo["dbCollection"].(string))
	_, err := lianjia.UpdateOne(ctx, bson.M{"Link": link}, bson.M{"$set": m})
	if err != nil {
		fmt.Print("update error!")
		fmt.Println(err)
	}
}

func AddZLItem(items []interface{}) {
	configInfo := configs.Config()
	client := GetInstance().client
	ctx := GetInstance().ctx

	db := client.Database(configInfo["dbDatabase"].(string))
	lianjia := db.Collection(configInfo["zlDBCollection"].(string))
	_, err := lianjia.InsertMany(ctx, items)
	if err != nil {
		if !strings.Contains(err.Error(), "multiple write errors") {
			fmt.Print("数据库插入失败！")
			fmt.Println(err)
		}
	}

}
