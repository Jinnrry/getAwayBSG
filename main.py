from spider import api

citylist = ['530', '601']
kw = ['php', 'java']
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
    print(item)

for cityid in citylist:
    for k in kw:
        start = 0
        total = 0
        res = api.getList(cityid, k, start, length)
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
                    }
                    saveData(data)
                if total > start + length:
                    res = api.getList(cityid, k, start, length)
                else:
                    break


