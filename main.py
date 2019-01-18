# coding:utf-8
from config.DBInfo import SessionFactory
from db.item import Item
from spider import api
from config import cityList

citylist = cityList.getCityList()

kw = ['php', 'java', 'python', 'c/c++', 'c#', 'mysql', 'oracle', 'javascript', 'linux', 'SQL', '软件', '程序员']
length = 50


def avgSalary(info):
    infos = info.split('-')
    total = 0
    for a in infos:
        a = a.replace('k', '000')
        a = a.replace('K', '000')
        a = a.replace('W', '0000')
        a = a.replace('w', '0000')
        try:
            a = int(a)
        except:
            a = 0
        total += a
    return (int)(total / len(infos))


def saveData(item):
    # 创建session对象:
    session = SessionFactory()
    # 创建新User对象:
    new_data = Item(item)
    # 添加到session:
    session.add(new_data)
    # 提交即保存到数据库:
    try:
        session.commit()
    except Exception as e:
        #print(e)
        pass


    # 关闭session:
    session.close()


for city in citylist:
    print(city['name'])
    for k in kw:
        start = 0
        total = 0
        res = api.getList(city['code'], k, start, length)
        while True:
            if res['code'] == 200:
                total = res['data']['numTotal']
                start += length
                for item in res['data']['results']:
                    data = {
                        'zlid': item['SOU_POSITION_ID'],
                        'score': item['score'],
                        'workingexp': item['workingExp']['name'],
                        'companyname': item['company']['name'],
                        'companysize': item['company']['size']['name'],
                        'companytype': item['company']['type']['name'],
                        'jobtype': item['jobType']['display'],
                        'createdate': item['createDate'],
                        'jobname': item['jobName'],
                        'enddate': item['endDate'],
                        'edulevel': item['eduLevel']['name'],
                        'city': item['city']['items'][0]['name'],
                        'salary': item['salary'],
                        'avgsalary': avgSalary(item['salary']),
                        'keyword': k,
                        'industry': 'it'
                    }
                    saveData(data)
                if total > start + length:
                    res = api.getList(city['code'], k, start, length)
                else:
                    break
