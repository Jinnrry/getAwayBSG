from config.DBInfo import SessionFactory
from db.ziroom import Ziroom
from spider import api

import requests
from PIL import Image
from pyquery import PyQuery as pq
import json
from urllib.request import urlretrieve
from tools import path

from pytesseract import pytesseract

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
    while True:
        html = content.html()
        start = html.find("ROOM_PRICE") + 12
        end = html.find("\n", start) - 1
        res = json.loads(html[start:end])
        numImage = 'http:' + res['image']
        filename = path.tempPath + 'numcode.png'
        urlretrieve(numImage, filename)
        text = pytesseract.image_to_string(Image.open(filename), config='-psm 7')
        text = text.replace(" ", "")
        if len(text) != 10:
            print("图片识别失败")
            continue
        offsets = res['offset']
        content = content('#houseList li').items()
        for index, house in enumerate(content):
            houseUrl = house('.txt a.t1').attr('href')
            offset = offsets[index]
            price = ''
            for i in offset:
                price += text[i:i + 1]
            # print(price)
            print(houseUrl)
            index_start = houseUrl.find("vr/")
            index_end = houseUrl.find(".html")
            houseid = houseUrl[index_start + 3:index_end]
            houseinfo = getHouseInfo(houseid, "110000")
            area = houseinfo['data']['area']
            iswhole = houseinfo['data']['is_whole']
            bedroom = houseinfo['data']['bedroom']
            parlor = houseinfo['data']['parlor']
            district_name = houseinfo['data']['district_name']
            bizcircle_name = houseinfo['data']['bizcircle_name']
            saveData({
                'id': houseid,
                'price': price,
                'url': 'http:' + houseUrl,
                'iswhole': iswhole,
                'bedroom': bedroom,
                'parlor': parlor,
                'district_name': district_name,
                'bizcircle_name': bizcircle_name,
                'area': area
            })
        break


def getHouseInfo(id, cityid):
    url = "http://m.ziroom.com/wap/detail/room.json?city_code=" + cityid + "&id=" + id
    res = api.get(url)
    res = json.loads(res.text)
    return res


def saveData(item):
    # 创建session对象:
    session = SessionFactory()
    # 创建新User对象:
    new_data = Ziroom(item)
    # 添加到session:
    session.add(new_data)
    # 提交即保存到数据库:
    try:
        session.commit()
    except Exception as e:
        if 'Duplicate' in repr(e):
            pass
            session.close()
            return False
        else:
            print(e)
            session.close()
            return False

    # 关闭session:
    session.close()
    return True


# http://m.ziroom.com/wap/detail/room.json?city_code=110000&id=61954203

getHouse(content)

# 2436851907
