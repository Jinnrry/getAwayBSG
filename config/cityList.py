import json
import os


def getCityList():
    with open(os.path.dirname(__file__) + "/cityList.json", 'r') as load_f:
        load_dict = json.load(load_f)
        return load_dict['cityList']
