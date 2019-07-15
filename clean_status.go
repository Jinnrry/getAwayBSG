package main

import (
	"context"
	"fmt"
	"getAwayBSG/configs"
	"getAwayBSG/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"
)

func main() {

	var choice int
	if len(os.Args) > 1 && strings.Index(os.Args[1], "lianjia") > -1 {
		choice = 1
	} else if len(os.Args) > 1 && strings.Index(os.Args[1], "zhilian") > -1 {
		choice = 2
	} else {
		fmt.Println("清除抓取状态（不清除状态的话爬虫会从上次停止位置继续抓取）")
		fmt.Println("请选择需要清除哪个爬虫的的状态数据：（输入数字）")
		fmt.Println("1.链家")
		fmt.Println("2.智联")
		fmt.Scanln(&choice)

	}

	if choice == 1 {
		db.SetLianjiaStatus(0)
		clean_visit()
		fmt.Println("Done!")
	} else if choice == 2 {
		db.SetZhilianStatus(0, 0)
		fmt.Println("Done!")
	} else {
		fmt.Println("选择错误！")
	}

}

func clean_visit() {
	conf := configs.Config()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := mongo.NewClient(options.Client().ApplyURI(conf["dburl"].(string) + "/colly"))
	if err := client.Connect(ctx); err != nil {
		fmt.Println(err)
	}

	odb := client.Database("colly")
	cookies := odb.Collection("colly_cookies")
	visit := odb.Collection("colly_visited")
	//清除全部的cookies
	_, err := cookies.DeleteMany(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = visit.DeleteMany(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

}