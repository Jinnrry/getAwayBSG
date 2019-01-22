#coding:utf-8
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


    def __init__(self,data):
        for key in data.keys():
            if key == 'id':
                self.id=data[key]
            if key == 'transactiondate':
                self.transactiondate=data[key]
            if key == 'price':
                self.price=data[key]
            if key == 'avgPrice':
                self.avgPrice=data[key]
            if key == 'ljID':
                self.ljID=data[key]
            if key == 'address':
                self.address=data[key]
