from spider import api
from bs4 import BeautifulSoup
res=api.get('https://www.lianjia.com/city/')
print(res.text)