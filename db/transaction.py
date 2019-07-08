# coding:utf-8
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, String, Integer, Float, DateTime

# 创建对象的基类:
Base = declarative_base()


class lianjia_transaction(Base):
    # 表的名字:
    __tablename__ = 'lianjia_transaction'

    # 表的结构:
    id = Column(Integer, primary_key=True)
    transactiondate = Column(DateTime)
    price = Column(Float)
    avgPrice = Column(Float)
    ljID = Column(Float)
    address = Column(String(255))
    address1 = Column(String(15))
    address2 = Column(String(15))
    address3 = Column(String(15))
    address4 = Column(String(15))
    address5 = Column(String(15))
    address6 = Column(String(15))
    address7 = Column(String(15))
    address8 = Column(String(15))
    address9 = Column(String(15))
    address10 = Column(String(15))
    url=Column(String(500))
    def __init__(self, data):
        for key in data.keys():
            if key == 'id':
                self.id = data[key]
            if key == 'transactiondate':
                self.transactiondate = data[key]
            if key == 'price':
                self.price = data[key]
            if key == 'avgPrice':
                self.avgPrice = data[key]
            if key == 'ljID':
                self.ljID = data[key]
            if key == 'address':
                self.address = data[key]
            if key == 'address1':
                self.address1 = data[key]
            if key == 'address2':
                self.address2 = data[key]
            if key == 'address3':
                self.address3 = data[key]
            if key == 'address4':
                self.address4 = data[key]
            if key == 'address5':
                self.address5 = data[key]
            if key == 'address6':
                self.address6 = data[key]
            if key == 'address7':
                self.address7 = data[key]
            if key == 'address8':
                self.address8 = data[key]
            if key == 'address9':
                self.address9 = data[key]
            if key == 'address10':
                self.address10 = data[key]
            if key == 'url':
                self.url = data[key]
