#coding:utf-8
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, String, Integer, Float, DateTime


# 创建对象的基类:
Base = declarative_base()

class Ziroom(Base):
    # 表的名字:
    __tablename__ = 'ziroom'

    # 表的结构:
    id = Column(Integer, primary_key=True)
    price = Column(Integer())
    url = Column(String(255))
    iswhole = Column(Integer())
    area = Column(Float())
    bedroom = Column(String(2))
    parlor = Column(String(2))
    district_name = Column(String(15))
    bizcircle_name = Column(String(15))


    def __init__(self,data):
        for key in data.keys():
            if key == 'id':
                self.id=data[key]
            if key == 'price':
                self.price=data[key]
            if key == 'url':
                self.url=data[key]
            if key == 'iswhole':
                self.iswhole=data[key]
            if key == 'area':
                self.area=data[key]
            if key == 'bedroom':
                self.bedroom=data[key]
            if key == 'parlor':
                self.parlor=data[key]
            if key == 'district_name':
                self.district_name=data[key]
            if key == 'bizcircle_name':
                self.bizcircle_name=data[key]
