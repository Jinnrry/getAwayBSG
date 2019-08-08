package entrance

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
	"regexp"
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

	if configInfo["crawlDelay"] != nil {
		delay, _ := configInfo["crawlDelay"].(json.Number).Int64()
		if delay > 0 {
			c.Limit(&colly.LimitRule{
				DomainGlob: "*",
				Delay:      time.Duration(delay) * time.Second,
			})
		}
	}

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

	c.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})

	c.OnHTML("body", func(element *colly.HTMLElement) {
		// 获取一页的数据
		element.ForEach(".LOGCLICKDATA", func(i int, e *colly.HTMLElement) {
			link := e.ChildAttr("a", "href")

			title := e.ChildText("a:first-child")
			fmt.Println(title)

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
			db.Add(bson.M{"zq_detail_status": 0, "Title": title, "TotalePrice": iPrice, "UnitPrice": iUnitPrice, "Link": link, "listCrawlTime": time.Now()})

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
			re, _ := regexp.Compile("pg\\d+/*")
			goUrl = re.ReplaceAllString(goUrl, "")
			err = c.Visit(goUrl)
			fmt.Println(err)
		})

		// 下一页
		element.ForEach(".page-box", func(i int, element *colly.HTMLElement) {
			var page Page
			err := json.Unmarshal([]byte(element.ChildAttr(".house-lst-page-box", "page-data")), &page)
			if err == nil {
				if page.CurPage < page.TotalPage {
					var gourl string
					re, _ := regexp.Compile("pg\\d+/*")
					gourl = re.ReplaceAllString(element.Request.URL.String(), "")
					gourl = gourl + "pg" + strconv.Itoa(page.CurPage+1)
					err = c.Visit(gourl)
					fmt.Println(err)
				}
			}
		})

	})

	err := c.Visit(cityUrl)
	fmt.Println(err)

}

func listCrawler() {
	confInfo := configs.Config()
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

	//设置延时
	if configInfo["crawlDelay"] != nil {
		delay, _ := configInfo["crawlDelay"].(json.Number).Int64()
		if delay > 0 {
			c.Limit(&colly.LimitRule{
				DomainGlob: "*",
				Delay:      time.Duration(delay) * time.Second,
			})
		}
	}

	//设置代理
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

	//随机UA
	extensions.RandomUserAgent(c)
	//自动referer
	extensions.Referer(c)
	//设置MongoDB存储状态信息
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

	c.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
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

			db.Update(element.Request.URL.String(), bson.M{"zq_detail_status": 1, "guapaitime": ttime, "detailCrawlTime": time.Now()})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("详情抓取：", r.URL.String())
	})

	client := db.GetClient()
	ctx := db.GetCtx()

	odb := client.Database(configInfo["dbDatabase"].(string))
	lianjia := odb.Collection(configInfo["dbCollection"].(string))

	//读取出全部需要抓取详情的数据
	cur, err := lianjia.Find(ctx, bson.M{"zq_detail_status": 0})

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

func Start_lianjia_ershou() {
	listFlag := make(chan int)   //记录列表抓取是否完成
	detailFlag := make(chan int) //记录详情是否抓取完成

	go func() {
		listCrawler()
		listFlag <- 1 //列表抓取完成
	}()

	go func() {
		zeroNum := 0
		for i := 0; i < 1; i = 0 {
			if crawlDetail() == 0 {
				zeroNum++
				if zeroNum > 3 { //尝试3次都没有详情需要抓取，结束详情抓取
					break
				}
				time.Sleep(300 * time.Second) //没有详情需要抓取了，等待5分钟再尝试
			}
		}
		detailFlag <- 1 //详情抓取完成
	}()

	//详情抓取与列表抓取都完成了，结束主线程
	<-listFlag
	<-detailFlag
}
