package db

import (
	"fmt"
	"getAwayBSG/configs"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
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
		fmt.Print("数据库更新出错！")
		fmt.Println(err)
	}
}

func AddZLItem(items []interface{}) {

	for i := 0; i < len(items); i++ {
		salary := items[i].(map[string]interface{})["salary"]
		if salary != nil && salary != "薪资面议" {

			K := strings.Index(salary.(string), "K")
			k := strings.Index(salary.(string), "k")
			Q := strings.Index(salary.(string), "千")

			W := strings.Index(salary.(string), "W")
			w := strings.Index(salary.(string), "w")
			Wan := strings.Index(salary.(string), "万")

			xishu := 0.0

			if K > 0 || k > 0 || Q > 0 {
				xishu = 1000
			} else if W > 0 || w > 0 || Wan > 0 {
				xishu = 10000
			} else {
				xishu = 1
			}

			salary = strings.Replace(salary.(string), "K", "", 2)
			salary = strings.Replace(salary.(string), "W", "", 2)
			salary = strings.Replace(salary.(string), "千", "", 2)
			salary = strings.Replace(salary.(string), "万", "", 2)
			salary = strings.Replace(salary.(string), "以下", "", 2)
			minAndMax := strings.Split(salary.(string), "-")

			min, err := strconv.ParseFloat(minAndMax[0], 32)
			if err != nil {
				min = 0
			}
			var max float64
			if len(minAndMax) > 1 {
				max, err = strconv.ParseFloat(minAndMax[1], 32)
				if err != nil {
					max = 0
				}
			} else {
				max = min
			}

			min = min * xishu
			max = max * xishu
			avg := (min + max) / 2

			items[i].(map[string]interface{})["min"] = min
			items[i].(map[string]interface{})["max"] = max
			items[i].(map[string]interface{})["avg"] = avg

		}

	}

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
