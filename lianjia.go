package main

import (
	"encoding/json"
	"fmt"
	"getAwayBSG/configs"
	"getAwayBSG/db"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	cachemongo "github.com/zolamk/colly-mongo-storage/colly/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	TotalPage int
	CurPage   int
}

func crawlerOneCity(cityUrl string) {
	c := colly.NewCollector()
	configInfo := configs.Config()
	if configInfo["proxyList"] != nil && len(configInfo["proxyList"].([]interface{})) > 0 {
		var proxyList []string
		for _, v := range configInfo["proxyList"].([]interface{}) {
			proxyList = append(proxyList, v.(string))
		}

		if configInfo["proxyList"] != nil {
			rp, err := proxy.RoundRobinProxySwitcher(proxyList...)
			if err != nil {
				fmt.Println(err)
			}
			c.SetProxyFunc(rp)
		}
	}
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	storage := &cachemongo.Storage{
		Database: "colly",
		URI:      configInfo["dburl"].(string) + "/colly",
	}

	if err := c.SetStorage(storage); err != nil {
		panic(err)
	}
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("列表抓取：", r.URL.String())
	})

	c.OnHTML("body", func(element *colly.HTMLElement) {
		// 获取一页的数据
		element.ForEach(".LOGCLICKDATA", func(i int, e *colly.HTMLElement) {
			link := e.ChildAttr("a", "href")

			title := e.ChildText("a:first-child")
			//fmt.Println(title)

			price := e.ChildText(".totalPrice")
			price = strings.Replace(price, "万", "0000", 1)
			//fmt.Println("总价：" + price)
			iPrice, err := strconv.Atoi(price)
			if err != nil {
				iPrice = 0
			}

			unitPrice := e.ChildAttr(".unitPrice", "data-price")

			//fmt.Println("每平米：" + unitPrice)
			//fmt.Println(e.Text)

			iUnitPrice, err := strconv.Atoi(unitPrice)
			if err != nil {
				iUnitPrice = 0
			}
			db.Add(bson.M{"Title": title, "TotalePrice": iPrice, "UnitPrice": iUnitPrice, "Link": link, "listCrawlTime": time.Now()})

		})

		// 切换地点
		element.ForEach(".position a", func(i int, element *colly.HTMLElement) {
			u, err := url.Parse(cityUrl)
			if err != nil {
				panic(err)
			}
			rootUrl := u.Scheme + "://" + u.Host

			goUrl := element.Attr("href")
			u, err = url.Parse(goUrl)
			if err != nil {
				fmt.Println(err)
			}
			if u.Scheme == "" {
				goUrl = rootUrl + u.Path
			} else {
				goUrl = u.String()
			}
			c.Visit(goUrl)
		})

		// 下一页
		element.ForEach(".page-box", func(i int, element *colly.HTMLElement) {
			var page Page
			json.Unmarshal([]byte(element.ChildAttr(".house-lst-page-box", "page-data")), &page)
			if page.CurPage < page.TotalPage {
				c.Visit(cityUrl + "pg" + strconv.Itoa(page.CurPage+1) + "/")
			}

		})

	})

	c.Visit(cityUrl)

}

func listCrawler() {
	confInfo := configs.Config()
	fmt.Print(confInfo)
	cityList := confInfo["cityList"].([]interface{})
	for i := db.GetLianjiaStatus(); i < len(cityList); i++ {
		crawlerOneCity(cityList[i].(string))
		db.SetLianjiaStatus(i)
	}
}

func crawlDetail() (sucnum int) {
	sucnum = 0
	c := colly.NewCollector()
	configInfo := configs.Config()

	if configInfo["proxyList"] != nil && len(configInfo["proxyList"].([]interface{})) > 0 {
		var proxyList []string
		for _, v := range configInfo["proxyList"].([]interface{}) {
			proxyList = append(proxyList, v.(string))
		}

		if configInfo["proxyList"] != nil {
			rp, err := proxy.RoundRobinProxySwitcher(proxyList...)
			if err != nil {
				fmt.Println(err)
			}
			c.SetProxyFunc(rp)
		}
	}

	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	storage := &cachemongo.Storage{
		Database: "colly",
		URI:      configInfo["dburl"].(string) + "/colly",
	}
	if err := c.SetStorage(storage); err != nil {
		panic(err)
	}
	c.OnHTML(".area .mainInfo", func(element *colly.HTMLElement) {
		area := strings.Replace(element.Text, "平米", "", 1)
		iArea, err := strconv.Atoi(area)
		if err != nil {
			iArea = 0
		}

		db.Update(element.Request.URL.String(), bson.M{"area": iArea, "detailCrawlTime": time.Now()})

	})

	c.OnHTML(".aroundInfo .communityName .info", func(element *colly.HTMLElement) {
		db.Update(element.Request.URL.String(), bson.M{"xiaoqu": element.Text, "detailCrawlTime": time.Now()})
	})

	c.OnHTML(".l-txt", func(element *colly.HTMLElement) {
		res := strings.Replace(element.Text, "二手房", "", 99)
		res = strings.Replace(res, " ", "", 99)
		address := strings.Split(res, ">")
		db.Update(element.Request.URL.String(), bson.M{"address": address[1 : len(address)-1], "detailCrawlTime": time.Now()})
	})

	c.OnHTML(".transaction li", func(element *colly.HTMLElement) {
		if element.ChildText("span:first-child") == "挂牌时间" {

			sGTime := element.ChildText("span:last-child")
			ttime, err := time.Parse("2006-01-02", sGTime)

			if err != nil {
				ttime = time.Now()
			}

			db.Update(element.Request.URL.String(), bson.M{"guapaitime": ttime, "detailCrawlTime": time.Now()})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("详情抓取：", r.URL.String())
	})

	client := db.GetClient()
	ctx := db.GetCtx()

	odb := client.Database(configInfo["dbDatabase"].(string))
	lianjia := odb.Collection(configInfo["dbCollection"].(string))

	cur, err := lianjia.Find(ctx, bson.M{"detailCrawlTime": bson.M{"$exists": false}})

	if err != nil {
		fmt.Println(err)
	} else {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var item bson.M
			if err := cur.Decode(&item); err != nil {
				fmt.Print("数据库读取失败！")
				fmt.Println(err)
			} else {
				sucnum++
				c.Visit(item["Link"].(string))
			}

		}
	}

	return sucnum
}

func main() {
	listFlag := make(chan int)
	detailFlag := make(chan int)

	go func() {
		listCrawler()
		listFlag <- 1
	}()

	go func() {
		zeroNum := 0
		for i := 0; i < 1; i = 0 {
			if crawlDetail() == 0 {
				zeroNum++
				if zeroNum > 3 {
					break
				}
				time.Sleep(300 * time.Second)
			}
		}
		detailFlag <- 1
	}()

	<-listFlag
	<-detailFlag
}
