import scrapy
from scrapy import Request


class CityListSpider(scrapy.Spider):
    name = "citylist"
    allowed_domains = ["lianjia.com"]
    # start_urls = [
    #     "https://www.lianjia.com/city/"
    # ]
    #
    # def parse(self, response):
    #     hreflist=response.selector.css(".city_list_ul a::attr(href)").extract()
    #     for url in hreflist:
    #         print(url)
    #         yield Request(url, callback=self.mainPage)
    #
    # def mainPage(self,response):
    #     print(response.url)
    #     title=response.xpath('//title/text()').extract()
    #     print(title)
    start_urls=['https://bj.lianjia.com/']

    def parse(self, response):
        print(response.body)
        print(response.status)