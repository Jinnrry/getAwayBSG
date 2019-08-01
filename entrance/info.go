package entrance

import (
	"getAwayBSG/configs"
	"getAwayBSG/db"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strconv"
	"strings"
	"time"
)

func Start_info(path string) {

	fd, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd_time := time.Now().Format("2006-01-02 15:04:05")
	fd_content := strings.Join([]string{
		fd_time, ":",
		getLianjiaErShouFangStatus(), " ",
		getLianJiaZuFangStatus(), " ",
		getZhiLianStatus(), "\n",
	}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
	fd.Close()

}

func getLianjiaErShouFangStatus() string {
	configInfo := configs.Config()
	client := db.GetClient()
	ctx := db.GetCtx()

	odb := client.Database(configInfo["dbDatabase"].(string))
	lianjia := odb.Collection(configInfo["dbCollection"].(string))
	lianjia_status := odb.Collection("lianjia_status")
	var info bson.M

	res := lianjia_status.FindOne(ctx, bson.M{})
	res.Decode(&info)
	detailNum, _ := lianjia.CountDocuments(ctx, bson.M{"address": bson.M{"$exists": true}})
	allNum, _ := lianjia.CountDocuments(ctx, bson.M{})

	return "链家二手房：详情数" + strconv.Itoa(int(detailNum)) + "总数：" + strconv.Itoa(int(allNum)) + " index:" + strconv.Itoa(int(info["index"].(int32)));
}

func getZhiLianStatus() string {
	configInfo := configs.Config()
	client := db.GetClient()
	ctx := db.GetCtx()

	odb := client.Database(configInfo["dbDatabase"].(string))
	zhilian := odb.Collection(configInfo["zlDBCollection"].(string))
	zhilian_status := odb.Collection("zhilian_status")

	zhilianNum, _ := zhilian.CountDocuments(ctx, bson.M{})
	var info bson.M
	res := zhilian_status.FindOne(ctx, bson.M{})
	res.Decode(&info)

	return "智联总数:" + strconv.Itoa(int(zhilianNum)) + " city_index:" + strconv.Itoa(int(info["city_index"].(int32))) + " kw_index:" + strconv.Itoa(int(info["kw_index"].(int32)))
}

func getLianJiaZuFangStatus() string {
	configInfo := configs.Config()
	client := db.GetClient()
	ctx := db.GetCtx()

	odb := client.Database(configInfo["dbDatabase"].(string))
	lianjiaZf := odb.Collection(configInfo["zufangCollection"].(string))
	lianjiaZFStatus := odb.Collection("lianjiazf_status")

	var info bson.M

	res := lianjiaZFStatus.FindOne(ctx, bson.M{})

	res.Decode(&info)

	allNum, _ := lianjiaZf.CountDocuments(ctx, bson.M{})

	return "链家租房：总数" + strconv.Itoa(int(allNum)) + " index:" + strconv.Itoa(int(info["index"].(int32)))
}
