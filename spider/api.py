import json

import requests


def get(url):
    return requests.get(url)


def getList(cityid, kw, start, length):
    url = 'https://fe-api.zhaopin.com/c/i/sou?start=' + str(start) + 'pageSize=' + str(length) + '&cityId=' + str(
        cityid) + '&workExperience=-1&education=-1&companyType=-1&employmentType=-1&jobWelfareTag=-1&sortType=publish&kw=' + str(
        kw) + '&kt=3&_v=0.17996222&x-zp-page-request-id=e8d2c03d3c4347a9b5edffa03367d90d-1547646999572-254944'
    return json.loads(get(url).text)
