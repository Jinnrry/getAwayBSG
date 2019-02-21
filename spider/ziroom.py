import requests
from pyquery import PyQuery as pq

header = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36'
}

response = requests.get("http://www.ziroom.com/z/nl/z3.html", headers=header)
content = pq(response.text)


def getEntrance(content):
    content = content('.clearfix.zIndex6')
    lista = content('a').items()
    for a in lista:
        url = a.attr('href')
        url = url.replace('//', 'http://', 1)
        print(url)


def getHouse(content):
    content = content('#houseList li').items()
    for house in content:
        houseUrl=house('.txt a.t1').attr('href')
        print(houseUrl)


getHouse(content)

#2436851907