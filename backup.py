import pymysql
import datetime
from config import DBInfo

# 打开数据库连接
db = pymysql.connect(DBInfo.dbhost, DBInfo.user, DBInfo.pwd, DBInfo.db)

# 使用 cursor() 方法创建一个游标对象 cursor
cursor = db.cursor()

today = datetime.date.today()
# 使用 execute()  方法执行 SQL 查询 
cursor.execute("CREATE TABLE `近12月生活压力_" + str(today) + "` AS SELECT * FROM `近12月生活压力` WHERE gfyl is NOT null;")

# 关闭数据库连接
db.close()
