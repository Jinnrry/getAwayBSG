import json
import re
from urllib.parse import urlparse
import scrapy
from scrapy import Request

from config.DBInfo import SessionFactory
from db.transaction import lianjia_transaction
from tools import tools


class LianJiaSpider(scrapy.Spider):
    name = "lianjia"
    start_urls = [
        "https://www.lianjia.com/city/"
    ]

    def parse(self, response):
        tools.writeLog("lianjia", "start")

        hreflist = response.selector.css(".city_list_ul a::attr(href)").extract()
        for url in hreflist:
            yield Request(url + "chengjiao", callback=self.mainPage)

    def mainPage(self, response):
        title = response.xpath('//title/text()').extract()
        if "|成交查询" in title[0]:
            res=self.recursiveListUrl(response)
            for url in res:
                yield Request(url, callback=self.getList)

            yield Request(response.url + "pg1/", callback=self.getList)

    def getList(self, response):
        res = self.recursiveListUrl(response)
        for url in res:
            yield Request(url, callback=self.getList)
        infourl = response.selector.css('.listContent .title  a::attr(href)').extract()
        for url in infourl:
            yield Request(url, callback=self.detail)

        try:
            strpageinfo = response.selector.css('.page-box .house-lst-page-box ::attr(page-data)').extract()[0]
            pageinfo = json.loads(strpageinfo)
            cur = pageinfo['curPage']
            total = pageinfo['totalPage']

            ourl=response.url
            result = urlparse(ourl)
            if "pg" not in result.path:
                ourl+="pg1/"
                result = urlparse(ourl)

            if cur == 1:
                while cur < total:
                    cur += 1
                    res = re.sub(r'pg\d+', "pg" + str(cur), result.path)
                    res = result.scheme + "://" + result.netloc + res
                    yield Request(res, callback=self.getList)
        except:
            pass

    def detail(self, response):
        # 成交时间
        date = response.selector.css('.house-title .wrapper span ::text').extract()[0][0:-2]
        price = response.selector.css('.info.fr .price i ::text').extract()[0]
        avgPrice = response.selector.css('.info.fr .price b ::text').extract()[0]
        ljID = response.selector.css('.transaction .content li:first-child ::text').extract()[1]
        address = ''
        address1 = ''
        address2 = ''
        address3 = ''
        address4 = ''
        address5 = ''
        address6 = ''
        address7 = ''
        address8 = ''
        address9 = ''
        address10 = ''
        index = 1
        for i in response.selector.css('.deal-bread ::text').extract()[1:-1]:
            i = i.replace("二手房成交价格", "")
            address += i
            if i != '' and i != '>':
                if index == 1:
                    address1 = i
                if index == 2:
                    address2 = i
                if index == 3:
                    address3 = i
                if index == 4:
                    address4 = i
                if index == 5:
                    address5 = i
                if index == 6:
                    address6 = i
                if index == 7:
                    address7 = i
                if index == 8:
                    address8 = i
                if index == 9:
                    address9 = i
                if index == 10:
                    address10 = i
                index += 1

        data = lianjia_transaction({
            'transactiondate': date,
            'price': float(price) * 10000,
            'avgPrice': avgPrice,
            'ljID': ljID.strip(),
            'address': address,
            'address1': address1,
            'address2': address2,
            'address3': address3,
            'address4': address4,
            'address5': address5,
            'address6': address6,
            'address7': address7,
            'address8': address8,
            'address9': address9,
            'address10': address10,
            'url': response.url
        })

        session = SessionFactory()
        # 添加到session:
        session.add(data)
        # 提交即保存到数据库:
        try:
            session.commit()
        except Exception as e:
            if 'Duplicate' in repr(e):
                session.close()
            else:
                session.close()

        # 关闭session:
        session.close()

    def writelog(self, url):
        with open("log.txt", 'a') as f:
            f.write(url + "\n")

    def recursiveListUrl(self, response):
        host = urlparse(response.url)
        host = host.scheme + "://" + host.netloc
        areaList = response.selector.css('.position a::attr(href)').extract()
        ret = []
        for url in areaList:
            if "https://" in url:
                ret.append(url)
            else:
                ret.append(host + url)
        return ret