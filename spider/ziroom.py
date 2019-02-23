# coding:utf-8
from tools import path
from config.DBInfo import SessionFactory
from db.ziroom import Ziroom
from spider import api
from tools import tools
from PIL import Image
from pyquery import PyQuery as pq
import json
from urllib.request import urlretrieve
from pytesseract import pytesseract

# 已经抓取过的URL
all = set('http://www.ziroom.com/z/nl/z3.html')
# 即将抓取的URL
already = set()


def getEntrance(content):
    content = content('.clearfix.zIndex6')
    lista = content('a').items()
    for a in lista:
        url = a.attr('href')
        url = url.replace('//', 'http://', 1)
        if url not in all:
            all.add(url)
            already.add(url)


# 获取一页数据
def getOnePage(url):
    print(url)
    while True:
        response = api.get(url)
        page = pq(response.text)

        # 提取区域URL
        getEntrance(page)
        html = page.html()
        start = html.find("ROOM_PRICE") + 12
        end = html.find("\n", start) - 1
        res = json.loads(html[start:end])
        numImage = 'http:' + res['image']
        filename = path.tempPath + 'numcode.png'
        urlretrieve(numImage, filename)
        text = pytesseract.image_to_string(Image.open(filename),lang="fontyp", config='--psm 7 digits')
        text = text.replace(" ", "")
        text = text.replace("/", "")
        text = text.replace("°", "7")
        if len(text) != 10:
            tools.writeLog("图片识别失败", "图片识别失败", "ziroom.log")
            continue
        offsets = res['offset']
        content = page('#houseList li').items()
        for index, house in enumerate(content):
            houseUrl = house('.txt a.t1').attr('href')
            offset = offsets[index]
            price = ''
            for i in offset:
                price += text[i:i + 1]
            index_start = houseUrl.find("vr/")
            index_end = houseUrl.find(".html")
            houseid = houseUrl[index_start + 3:index_end]
            houseinfo = getHouseInfo(houseid, "110000")
            if houseinfo['error_code'] != 0:
                continue
            if houseinfo['data']['price_unit'] == "/天":
                price = int(price) * 30
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
    nextPage = page("#page .next").attr("href")
    if nextPage != None:
        nextPage = 'http:' + nextPage
        getOnePage(nextPage)


def getHouseInfo(id, cityid):
    url = "http://m.ziroom.com/wap/detail/room.json?city_code=" + cityid + "&id=" + id
    res = api.get(url)
    try:
        res = json.loads(res.text)
    except:
        return {'error_code': 500}
    return res


def saveData(item):
    session = SessionFactory()
    new_data = Ziroom(item)
    session.add(new_data)
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

    session.close()
    return True


def start():
    tools.writeLog("启动", "自如爬虫")
    getOnePage('http://www.ziroom.com/z/nl/z3.html')
    while True:
        try:
            url = already.pop()
            getOnePage(url)
        except:
            break


start()
