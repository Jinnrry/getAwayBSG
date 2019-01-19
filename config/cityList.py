#coding:utf-8
import json
import os

#设置抓取城市，修改json文件
def getCityList():
    with open(os.path.dirname(__file__) + "/cityList.json", 'r',encoding='UTF-8') as load_f:
        load_dict = json.load(load_f)
        return load_dict['cityList']
