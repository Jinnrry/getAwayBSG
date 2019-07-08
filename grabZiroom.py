#coding:utf-8
import json
import os
import smtplib
from email.mime.text import MIMEText

from spider import api

def getHouseInfo(id, cityid):
    url = "http://m.ziroom.com/wap/detail/room.json?city_code=" + cityid + "&id=" + id
    res = api.get(url)
    try:
        res = json.loads(res.text)
    except:
        return {'error_code': 500}
    return res

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

res=getHouseInfo('61411849',"110000")

if res['data']['status'] != 'tzpzz':
    if not os.path.exists('lock'):
        fp = open("lock",'w')
        fp.close()
        print("send!")
        sendEmail("自如房屋释放通知","自如房屋释放通知！房屋链接：http://www.ziroom.com/z/vr/61411849.html","ok@xjiangwei.cn")
        sendEmail("自如房屋释放通知","自如房屋释放通知！房屋链接：http://www.ziroom.com/z/vr/61411849.html","jiangwei@mafengwo.com")
        sendEmail("自如房屋释放通知","自如房屋释放通知！房屋链接：http://www.ziroom.com/z/vr/61411849.html","604836556@qq.com")
    else:
        exit()
