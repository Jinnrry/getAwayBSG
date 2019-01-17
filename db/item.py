from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, String, Integer, Float, DateTime


# 创建对象的基类:
Base = declarative_base()

class Item(Base):
    # 表的名字:
    __tablename__ = 'item'

    # 表的结构:
    id = Column(Integer, primary_key=True)
    zlid = Column(String(50))
    score = Column(Float)
    workingexp = Column(String(30))
    companyname = Column(String(100))
    companysize = Column(String(30))
    companytype = Column(String(15))
    jobtype = Column(String(200))
    createdate = Column(DateTime)
    jobname = Column(String(50))
    enddate = Column(DateTime)
    edulevel = Column(String(15))
    city = Column(String(15))
    salary = Column(String(20))
    avgsalary = Column(Integer)

    def __init__(self,data):
        # for key in data.keys():
        #     self[key]=data[key]

        print(data.keys())
