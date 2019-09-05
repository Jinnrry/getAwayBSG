package entrance

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"getAwayBSG/configs"
	"getAwayBSG/db"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Start_zhilian() {
	configInfo := configs.Config()
	keys := configInfo["zlKeyWords"].([]interface{})
	cityList := configInfo["zlCityList"].([]interface{})

	cityIndex, kwIndex := db.GetZhilianStatus()

	for i := kwIndex; i < len(keys); i++ {
		for j := cityIndex; j < len(cityList); j++ {
			var total int = 1000
			for start := 0; start < total; start += 50 {
				cityid := cityList[j].(map[string]interface{})["code"]
				if cityid == nil {
					fmt.Println(cityList[j])
				}
				icityid, err := cityid.(json.Number).Int64()
				if err != nil {
					icityid = 530
				}
				length := 50
				keyword := keys[i].(string)
				keyword = url.QueryEscape(keyword)
				////apiUrl:= "https://fe-api.zhaopin.com/c/i/sou?start=" + strconv.Itoa(start) + "pageSize=" + strconv.Itoa(length) + "&cityId=" + strconv.Itoa(cityid) + "&workExperience=-1&education=-1&companyType=-1&employmentType=-1&jobWelfareTag=-1&sortType=publish&kw=" + keys[i].(string) + "&kt=3&_v=0.17996222&x-zp-page-request-id=e8d2c03d3c4347a9b5edffa03367d90d-1547646999572-254944"
				apiUrl := "https://fe-api.zhaopin.com/c/i/sou?start=" + strconv.Itoa(start) + "&pageSize=" + strconv.Itoa(length) + "&cityId=" + strconv.Itoa(int(icityid)) + "&workExperience=-1&education=-1&companyType=-1&employmentType=-1&jobWelfareTag=-1&sortType=publish&kw=" + keyword + "&kt=3&_v=0.96788938&x-zp-page-request-id=adce992a71af4857ad9dd407cae222ff-1562161856663-558612&x-zp-client-id=f0fe8f7b-8a03-4076-9894-4389e9959954"
				fmt.Println(apiUrl)
				res := get(apiUrl)
				var mapResult map[string]interface{}
				err = json.Unmarshal([]byte(res), &mapResult)

				if err != nil {
					fmt.Println("JsonToMapDemo err: ", err)
				} else {
					if mapResult["data"] != nil {
						data := mapResult["data"].(map[string]interface{})
						numTotal := data["numTotal"]
						total = int(numTotal.(float64))
						results := data["results"].([]interface{})
						for index := range results {
							var itemTime string
							loc, _ := time.LoadLocation("Local")
							itemTime = results[index].(map[string]interface{})["updateDate"].(string)
							results[index].(map[string]interface{})["updateDate"], _ = time.ParseInLocation("2006-01-02 15:04:05", itemTime, loc)
							results[index].(map[string]interface{})["crawler_time"] = time.Now()
						}
						db.AddZLItem(results)
					} else {
						fmt.Println("接口返回错误！")
					}
				}
			}
			db.SetZhilianStatus(j, i)
		}
	}
	//fmt.Println(keys[i])

}

func get(link string) (bodystr string) {

	bodystr = ""
	var client *http.Client
	configInfo := configs.Config()

	if configInfo["crawlDelay"] != nil {
		delay, _ := configInfo["crawlDelay"].(json.Number).Int64()
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}

	if configInfo["proxyList"] != nil {
		var proxyList []string
		for _, v := range configInfo["proxyList"].([]interface{}) {
			proxyList = append(proxyList, v.(string))
		}
		rand.Seed(time.Now().Unix())

		proxy, _ := url.Parse(proxyList[rand.Intn(len(proxyList))])
		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client = &http.Client{
			Transport: tr,
			Timeout:   time.Second * 10, //超时时间
		}
	} else {
		client = &http.Client{
			Timeout: time.Second * 10, //超时时间
		}
	}

	reqest, _ := http.NewRequest("GET", link, nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	response, _ := client.Do(reqest)
	if response != nil {
		if response.StatusCode == 200 {
			body, _ := ioutil.ReadAll(response.Body)
			bodystr = string(body)
		}
	}

	return bodystr
}
