# 逃离北上广


[![Home](https://img.shields.io/badge/link-项目主页-brightgreen.svg)](https://jinnrry.github.io/getAwayBSG/)
[![Link](https://img.shields.io/badge/link-python实现-blue.svg)](https://github.com/jinnrry/getAwayBSG/tree/python)
[![Downloads](https://img.shields.io/github/downloads/jinnrry/getAwayBSG/total)](https://img.shields.io/github/downloads/jinnrry/getAwayBSG/total)
[![forks](https://img.shields.io/github/forks/jinnrry/getAwayBSG?style=flat)](https://img.shields.io/github/forks/jinnrry/getAwayBSG?style=flat)
[![starts](https://img.shields.io/github/stars/jinnrry/getAwayBSG)](https://img.shields.io/github/stars/jinnrry/getAwayBSG)
[![license](https://img.shields.io/github/license/jinnrry/getAwayBSG)](https://img.shields.io/github/license/jinnrry/getAwayBSG)
[![issues](https://img.shields.io/github/issues/jinnrry/getAwayBSG)](https://img.shields.io/github/issues/jinnrry/getAwayBSG)
[![version](https://img.shields.io/github/release/jinnrry/getAwayBSG)](https://img.shields.io/github/release/jinnrry/getAwayBSG)



> **注意！**\
> 1.本项目仅供学习研究，禁止用于任何商业项目\
> 2.运行的时候为被爬方考虑下！尽量不要爬全站。请在配置文件中设置你需要的城市爬取即可！\
> 3.[项目主页](https://jinnrry.github.io/getAwayBSG/)里面有现成数据，不需要你自己动手运行爬虫 


## 啥？

如果你是一个正准备逃离北上广等一线城市却又找不到去处的IT人士，或许这个项目能给你点建议。

通过爬虫，抓取了链接、智联的工作，租房，二手房一系列数据，为你提供各城市的宏观分析数据

## 安装

从[releases](https://github.com/jinnrry/getAwayBSG/releases)下载对应操作系统，对应平台的二进制文件和配置文件模板

## 配置

打开配置文件你就知道了

## 运行

链家二手房数据抓取

```
getAwayBSG -config=config.yaml -lianjia_ershou
```

链家租房数据抓取

```
getAwayBSG -config=config.yaml -lianjia_zufang
```

智联招聘数据抓取

```
getAwayBSG -config=config.yaml -zhilian
```

其他命令

1.clean

清除缓存状态，抓取过程中会将抓取进度保存到mongodb，每次启动会从上次位置继续抓取，如果你需要清除缓存状态，执行
```
getAwayBSG -clean 
```
该命令支持脚本调用
```
getAwayBSG -clean [option]
```

option支持：lianjia_ershou、zhilian、lianjia_zufang

2.info

方便定时脚本记录抓取情况，使用info命令可以输出当前抓取数据量到文件

```
getAwayBSG -info -info_save_to=./numLog.txt
```

使用-info_save_to参数指定文件保存位置，默认为当前目录的numLog.txt文件中

3.help

输出支持的全部命令列表

```
getAwayBSG -help
```


## 数据分析

分析用的MongoDB语句在[Query.js](./Query.js)文件中，使用MongoDB执行即可

## 编译

编译使用xgo，需要先安装docker

```
git clone https://github.com/jinnrry/getAwayBSG

docker pull karalabe/xgo-latest

go get github.com/karalabe/xgo

cd getAwayBSG

sh ./build.sh
```

## 部署

如果需要分布式或者多进程抓取，在不同机器或者多个进程中指定相同的MongoDB源即可，程序已经支持分布式多进程抓取了。已抓取的链接和状态会通过MongoDB共享
