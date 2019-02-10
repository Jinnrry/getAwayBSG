import pymysql
 
# 打开数据库连接
db = pymysql.connect("localhost","zhaopin","root","78667602" )
 
# 使用 cursor() 方法创建一个游标对象 cursor
cursor = db.cursor()
 
# 使用 execute()  方法执行 SQL 查询 
cursor.execute("


CREATE TABLE `近12月生活压力_` AS SELECT
    *
FROM
    `近12月生活压力` WHERE gfyl is NOT null;

")
 
 
# 关闭数据库连接
db.close()
