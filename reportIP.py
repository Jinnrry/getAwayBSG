#coding:utf-8

import requests
import smtplib
from email.mime.text import MIMEText

def sendEmail(title,content,user):
    mailserver = "smtp.163.com"  #邮箱服务器地址
    username_send = 'jiangwei1995910@163.com'  #邮箱用户名
    password = 'jw631106246'   #邮箱密码：需要使用授权码
    username_recv = user  #收件人，多个收件人用逗号隔开
    mail = MIMEText(content)
    mail['Subject'] = title
    mail['From'] = username_send  #发件人
    mail['To'] = username_recv  #收件人；[]里的三个是固定写法，别问为什么，我只是代码的搬运工
    smtp = smtplib.SMTP(mailserver,port=25) # 连接邮箱服务器，smtp的端口号是25
    # smtp=smtplib.SMTP_SSL('smtp.qq.com',port=465) #QQ邮箱的服务器和端口号
    smtp.login(username_send,password)  #登录邮箱
    smtp.sendmail(username_send,username_recv,mail.as_string())# 参数分别是发送者，接收者，第三个是把上面的发送邮件的内容变成字符串
    smtp.quit() # 发送完毕后退出smtp


headers = {'User-Agent': 'curl/7.54.0'}
res=requests.get("https://www.ip.cn",headers=headers)


sendEmail("定时脚本：汇报IP",res.text,"ok@xjiangwei.cn")
