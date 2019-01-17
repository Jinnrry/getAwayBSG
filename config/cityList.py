#coding:utf-8
import json
import os


def getCityList():
    with open(os.path.dirname(__file__) + "/cityList.json", 'r',encoding='UTF-8') as load_f:
        load_dict = json.load(load_f)
        return load_dict['cityList']
