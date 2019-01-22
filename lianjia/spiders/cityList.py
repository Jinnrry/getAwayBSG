import json

import scrapy
from scrapy import Request

from config.DBInfo import SessionFactory
from db.transaction import lianjia_transaction


class CityListSpider(scrapy.Spider):
    name = "citylist"
    start_urls = [
        "https://www.lianjia.com/city/"
    ]

    def parse(self, response):
        hreflist = response.selector.css(".city_list_ul a::attr(href)").extract()
        for url in hreflist:
            yield Request(url + "chengjiao", callback=self.mainPage)

    def mainPage(self, response):
        title = response.xpath('//title/text()').extract()

        if "|成交查询" in title[0]:
            yield Request(response.url + "pg1/", callback=self.getList)

    def getList(self, response):
        infourl = response.selector.css('.listContent .title  a::attr(href)').extract()
        for url in infourl:
            yield Request(url, callback=self.detail)

        strpageinfo = response.selector.css('.page-box .house-lst-page-box ::attr(page-data)').extract()[0]
        pageinfo = json.loads(strpageinfo)
        if pageinfo['curPage'] < pageinfo['totalPage']:
            detailUrl = self.buildUlr(response.url.split('/')) + "pg" + str(pageinfo['curPage'] + 1)
            yield Request(detailUrl, callback=self.getList)

    def buildUlr(self, args):
        ret = ''
        for item in args[0:4]:
            ret += item + "/"
        return ret

    def detail(self, response):
        # 成交时间
        date = response.selector.css('.house-title .wrapper span ::text').extract()[0][0:-2]
        price = response.selector.css('.info.fr .price i ::text').extract()[0]
        avgPrice = response.selector.css('.info.fr .price b ::text').extract()[0]
        ljID = response.selector.css('.transaction .content li:first-child ::text').extract()[1]
        address = ''
        for i in response.selector.css('.deal-bread ::text').extract()[1:-1]:
            address += i.replace("二手房成交价格","")
        data = lianjia_transaction({
            'transactiondate': date,
            'price': float(price) * 10000,
            'avgPrice': avgPrice,
            'ljID': ljID.strip(),
            'address': address
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
                print(e)
                session.close()

        # 关闭session:
        session.close()
