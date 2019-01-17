import json
def getCityList():
    with open("./cityList.json", 'r') as load_f:
        load_dict = json.load(load_f)
        return load_dict['cityList']
