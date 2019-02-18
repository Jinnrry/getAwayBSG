# coding:utf-8

import requests
import smtplib
from email.mime.text import MIMEText

import pymysql
import datetime


def getStatus():
    # 打开数据库连接
    db = pymysql.connect("localhost", "root", "78667602", "zhaopin")

    # 使用 cursor() 方法创建一个游标对象 cursor
    cursor = db.cursor()

    today = datetime.date.today()
    # 使用 execute()  方法执行 SQL 查询
    cursor.execute("select count(1) from item;")
    results = cursor.fetchall()
    itemNum = results[0][0]

    cursor.execute("select count(1) from lianjia_transaction;")
    results = cursor.fetchall()
    lianjiaNum = results[0][0]

    # 关闭数据库连接
    db.close()

    return {"lianjia": lianjiaNum, "zhilian": itemNum}


def sendEmail(title, content, user):
    mailserver = "smtp.163.com"  # 邮箱服务器地址
    username_send = 'jiangwei1995910@163.com'  # 邮箱用户名
    password = 'jw631106246'  # 邮箱密码：需要使用授权码
    username_recv = user  # 收件人，多个收件人用逗号隔开
    mail = MIMEText(content)
    mail['Subject'] = title
    mail['From'] = username_send  # 发件人
    mail['To'] = username_recv  # 收件人；[]里的三个是固定写法，别问为什么，我只是代码的搬运工
    smtp = smtplib.SMTP(mailserver, port=25)  # 连接邮箱服务器，smtp的端口号是25
    # smtp=smtplib.SMTP_SSL('smtp.qq.com',port=465) #QQ邮箱的服务器和端口号
    smtp.login(username_send, password)  # 登录邮箱
    smtp.sendmail(username_send, username_recv, mail.as_string())  # 参数分别是发送者，接收者，第三个是把上面的发送邮件的内容变成字符串
    smtp.quit()  # 发送完毕后退出smtp


def getIP():
    headers = {'User-Agent': 'curl/7.54.0'}
    res = requests.get("https://www.ip.cn", headers=headers)
    return res.text


def buildBody():
    ip = getIP()
    status = getStatus()
    return "服务器网络信息：" + ip + "<br>服务器数据信息：<br>智联数据量：" + status['zhilian'] + "链家数据量：" + status['lianjia']


sendEmail("定时脚本：汇报状态", buildBody(), "ok@xjiangwei.cn")
