CREATE TABLE `item` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `zlid` varchar(50) COLLATE utf8_bin NOT NULL COMMENT '智联ID',
 `score` double NOT NULL COMMENT '智联评分',
 `workingexp` varchar(30) COLLATE utf8_bin NOT NULL COMMENT '工作年限',
 `companyname` varchar(100) COLLATE utf8_bin NOT NULL COMMENT '公司名称',
 `companysize` varchar(30) COLLATE utf8_bin NOT NULL COMMENT '公司规模',
 `companytype` varchar(15) COLLATE utf8_bin NOT NULL COMMENT '公司类型',
 `jobtype` varchar(200) COLLATE utf8_bin NOT NULL COMMENT '工作类型',
 `createdate` datetime NOT NULL COMMENT '创建时间',
 `jobname` varchar(50) COLLATE utf8_bin NOT NULL COMMENT '工作名称',
 `enddate` datetime NOT NULL COMMENT '结束时间',
 `edulevel` varchar(15) COLLATE utf8_bin NOT NULL COMMENT '教育程度',
 `city` varchar(15) COLLATE utf8_bin NOT NULL COMMENT '城市',
 `salary` varchar(20) COLLATE utf8_bin NOT NULL COMMENT '薪资',
 `avgsalary` int(11) NOT NULL COMMENT '平均薪资',
 `zqtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '抓取时间',
 `keyword` varchar(20) COLLATE utf8_bin DEFAULT NULL,
 `industry` varchar(20) COLLATE utf8_bin DEFAULT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `zlid_index` (`zlid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;



CREATE TABLE `lianjia_transaction` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `transactiondate` datetime NOT NULL COMMENT '成交时间',
 `zqtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '抓取时间',
 `price` double NOT NULL COMMENT '成交价格',
 `avgPrice` double NOT NULL COMMENT '每平米价格',
 `ljID` varchar(25) NOT NULL COMMENT '链家ID',
 `address` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '地址',
 `address1` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address2` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address3` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address4` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address5` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address6` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address7` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address8` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address9` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `address10` varchar(15) COLLATE utf8_bin DEFAULT NULL,
 `url` varchar(500) COLLATE utf8_bin DEFAULT NULL,
 PRIMARY KEY (`id`),
 UNIQUE KEY `ljID` (`ljID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;



CREATE TABLE `ziroom` (
  `id` int(11) NOT NULL,
  `price` int(10) DEFAULT NULL,
  `url` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `iswhole` tinyint(1) DEFAULT NULL,
  `ctime` datetime DEFAULT CURRENT_TIMESTAMP,
  `area` varchar(10) DEFAULT NULL,
  `bedroom` varchar(2) COLLATE utf8_bin DEFAULT NULL,
  `parlor` varchar(2) COLLATE utf8_bin DEFAULT NULL,
  `district_name` varchar(15) COLLATE utf8_bin DEFAULT NULL,
  `bizcircle_name` varchar(15) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


--
-- 视图结构 `1-3年城市薪资`
--
DROP TABLE IF EXISTS `1-3年城市薪资`;

CREATE VIEW `1-3年城市薪资`  AS  select avg(`item`.`avgsalary`) AS `AVG(avgsalary)`,std(`item`.`avgsalary`) AS `STD(avgsalary)`,`item`.`city` AS `city`,count(1) AS `num` from `item` where (`item`.`workingexp` = '1-3年') group by `item`.`city` having (`num` > 300) order by `AVG(avgsalary)` desc ;

-- --------------------------------------------------------

--
-- 视图结构 `3-5年薪资`
--
DROP TABLE IF EXISTS `3-5年薪资`;

CREATE VIEW `3-5年薪资`  AS  select avg(`item`.`avgsalary`) AS `AVG(avgsalary)`,std(`item`.`avgsalary`) AS `STD(avgsalary)`,`item`.`city` AS `city`,count(1) AS `num` from `item` where (`item`.`workingexp` = '3-5年') group by `item`.`city` having (`num` > 300) order by `AVG(avgsalary)` desc ;

-- --------------------------------------------------------

--
-- 视图结构 `2015年数据`
--
DROP TABLE IF EXISTS `2015年数据`;

CREATE VIEW `2015年数据`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2` from `lianjia_transaction` where ((`lianjia_transaction`.`address1` = '成都') and (`lianjia_transaction`.`transactiondate` < '2016-01-01') and (`lianjia_transaction`.`transactiondate` > '2015-01-01')) group by `lianjia_transaction`.`address2` ;

-- --------------------------------------------------------

--
-- 视图结构 `2016年数据`
--
DROP TABLE IF EXISTS `2016年数据`;

CREATE VIEW `2016年数据`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2` from `lianjia_transaction` where ((`lianjia_transaction`.`address1` = '成都') and (`lianjia_transaction`.`transactiondate` < '2017-01-01') and (`lianjia_transaction`.`transactiondate` > '2016-01-01')) group by `lianjia_transaction`.`address2` ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年数据`
--
DROP TABLE IF EXISTS `2017年数据`;

CREATE VIEW `2017年数据`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2` from `lianjia_transaction` where ((`lianjia_transaction`.`address1` = '成都') and (`lianjia_transaction`.`transactiondate` < '2018-01-01') and (`lianjia_transaction`.`transactiondate` > '2017-01-01')) group by `lianjia_transaction`.`address2` ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年起价格走势`
--
DROP TABLE IF EXISTS `2017年起价格走势`;

CREATE VIEW `2017年起价格走势`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2`,count(1) AS `COUNT(1)`,date_format(`lianjia_transaction`.`transactiondate`,'%Y%m') AS `tdate` from `lianjia_transaction` where (`lianjia_transaction`.`transactiondate` > '2017-01-01') group by `tdate`,`lianjia_transaction`.`address2` order by `lianjia_transaction`.`address2`,`tdate` ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年起全国价格走势`
--
DROP TABLE IF EXISTS `2017年起全国价格走势`;

CREATE VIEW `2017年起全国价格走势`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,count(1) AS `COUNT(1)`,date_format(`lianjia_transaction`.`transactiondate`,'%Y%m') AS `tdate` from `lianjia_transaction` where (`lianjia_transaction`.`transactiondate` > '2017-01-01') group by `tdate` order by `tdate` ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年起北京价格走势`
--
DROP TABLE IF EXISTS `2017年起北京价格走势`;

CREATE VIEW `2017年起北京价格走势`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2`,count(1) AS `COUNT(1)`,date_format(`lianjia_transaction`.`transactiondate`,'%Y%m') AS `tdate` from `lianjia_transaction` where ((`lianjia_transaction`.`transactiondate` > '2017-01-01') and (`lianjia_transaction`.`address1` = '北京')) group by `tdate`,`lianjia_transaction`.`address2` order by `lianjia_transaction`.`address2`,`tdate` ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年起各城市房屋均价`
--
DROP TABLE IF EXISTS `2017年起各城市房屋均价`;

CREATE VIEW `2017年起各城市房屋均价`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `AVG(avgPrice)`,avg(`lianjia_transaction`.`price`) AS `AVG(price)`,`lianjia_transaction`.`address1` AS `address1` from `lianjia_transaction` where (`lianjia_transaction`.`transactiondate` > '2017-01-01') group by `lianjia_transaction`.`address1` order by `AVG(price)` desc ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年起成都价格走势`
--
DROP TABLE IF EXISTS `2017年起成都价格走势`;

CREATE VIEW `2017年起成都价格走势`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2`,count(1) AS `COUNT(1)`,date_format(`lianjia_transaction`.`transactiondate`,'%Y%m') AS `tdate` from `lianjia_transaction` where ((`lianjia_transaction`.`transactiondate` > '2017-01-01') and (`lianjia_transaction`.`address1` = '成都')) group by `tdate`,`lianjia_transaction`.`address2` order by `lianjia_transaction`.`address2`,`tdate` ;

-- --------------------------------------------------------

--
-- 视图结构 `2017年起深圳价格走势`
--
DROP TABLE IF EXISTS `2017年起深圳价格走势`;

CREATE VIEW `2017年起深圳价格走势`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `avg(``avgPrice``)`,std(`lianjia_transaction`.`avgPrice`) AS `std(``avgPrice``)`,avg(`lianjia_transaction`.`price`) AS `avg(``price``)`,std(`lianjia_transaction`.`price`) AS `std(``price``)`,`lianjia_transaction`.`address2` AS `address2`,count(1) AS `COUNT(1)`,date_format(`lianjia_transaction`.`transactiondate`,'%Y%m') AS `tdate` from `lianjia_transaction` where ((`lianjia_transaction`.`transactiondate` > '2017-01-01') and (`lianjia_transaction`.`address1` = '深圳')) group by `tdate`,`lianjia_transaction`.`address2` order by `lianjia_transaction`.`address2`,`tdate` ;

-- --------------------------------------------------------

--
-- 视图结构 `20180101开始成都数据`
--
DROP TABLE IF EXISTS `20180101开始成都数据`;

CREATE VIEW `20180101开始成都数据`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `AVG(avgPrice)`,std(`lianjia_transaction`.`avgPrice`) AS `std(avgPrice)`,avg(`lianjia_transaction`.`price`) AS `AVG(price)`,std(`lianjia_transaction`.`price`) AS `STD(price)`,`lianjia_transaction`.`address2` AS `address2` from `lianjia_transaction` where ((`lianjia_transaction`.`address1` = '成都') and (`lianjia_transaction`.`transactiondate` > '2018-01-01')) group by `lianjia_transaction`.`address2` ;

-- --------------------------------------------------------

--
-- 视图结构 `citySarly`
--
DROP TABLE IF EXISTS `citySarly`;

CREATE VIEW `citySarly`  AS  select avg(`item`.`avgsalary`) AS `avgsalarys`,std(`item`.`avgsalary`) AS `STD(avgsalary)`,`item`.`city` AS `city`,count(1) AS `counts` from `item` where (`item`.`avgsalary` <> 0) group by `item`.`city` having (`counts` > 300) order by `avgsalarys` desc ;

-- --------------------------------------------------------

--
-- 视图结构 `城市薪资`
--
DROP TABLE IF EXISTS `城市薪资`;

CREATE VIEW `城市薪资`  AS  select avg(`item`.`avgsalary`) AS `AVG(avgsalary)`,std(`item`.`avgsalary`) AS `STD(avgsalary)`,`item`.`city` AS `city`,count(1) AS `num` from `item` group by `item`.`city` having (`num` > 300) order by `AVG(avgsalary)` desc ;

-- --------------------------------------------------------

--
-- 视图结构 `按月统计薪资`
--
DROP TABLE IF EXISTS `按月统计薪资`;

CREATE VIEW `按月统计薪资`  AS  select avg(`a`.`avgsalary`) AS `AVG(avgsalary)`,count(1) AS `COUNT(1)`,`a`.`DATE` AS `DATE` from (select `item`.`id` AS `id`,`item`.`zlid` AS `zlid`,`item`.`zqtime` AS `zqtime`,`item`.`score` AS `score`,`item`.`workingexp` AS `workingexp`,`item`.`companyname` AS `companyname`,`item`.`companysize` AS `companysize`,`item`.`companytype` AS `companytype`,`item`.`jobtype` AS `jobtype`,`item`.`createdate` AS `createdate`,`item`.`jobname` AS `jobname`,`item`.`enddate` AS `enddate`,`item`.`edulevel` AS `edulevel`,`item`.`city` AS `city`,`item`.`salary` AS `salary`,`item`.`avgsalary` AS `avgsalary`,`item`.`keyword` AS `keyword`,`item`.`industry` AS `industry`,date_format(`item`.`createdate`,'%Y%m') AS `DATE` from `item` where (`item`.`avgsalary` <> 0)) `a` group by `a`.`DATE` ;

-- --------------------------------------------------------

--
-- 视图结构 `生活压力`
--
DROP TABLE IF EXISTS `生活压力`;

CREATE VIEW `生活压力`  AS  select `a`.`avgsalarys` AS `avgsalarys`,`a`.`STD(avgsalary)` AS `STD(avgsalary)`,`a`.`city` AS `city`,`a`.`counts` AS `counts`,`b`.`AVG(avgPrice)` AS `AVG(avgPrice)`,`b`.`AVG(price)` AS `AVG(price)`,`b`.`address1` AS `address1`,(`a`.`avgsalarys` / `b`.`AVG(avgPrice)`) AS `gfyl` from (`citySarly` `a` left join `2017年起各城市房屋均价` `b` on((`a`.`city` = `b`.`address1`))) order by (`a`.`avgsalarys` / `b`.`AVG(avgPrice)`) desc ;

-- --------------------------------------------------------

--
-- 视图结构 `近12月城市房价`
--
DROP TABLE IF EXISTS `近12月城市房价`;

CREATE VIEW `近12月城市房价`  AS  select avg(`lianjia_transaction`.`avgPrice`) AS `AVG(avgPrice)`,avg(`lianjia_transaction`.`price`) AS `AVG(price)`,`lianjia_transaction`.`address1` AS `address1` from `lianjia_transaction` where (`lianjia_transaction`.`transactiondate` between (now() - interval 12 month) and now()) group by `lianjia_transaction`.`address1` order by `AVG(price)` desc ;

-- --------------------------------------------------------

--
-- 视图结构 `近12月生活压力`
--
DROP TABLE IF EXISTS `近12月生活压力`;

CREATE VIEW `近12月生活压力`  AS  select `a`.`avgsalarys` AS `avgsalarys`,`a`.`STD(avgsalary)` AS `STD(avgsalary)`,`a`.`city` AS `city`,`a`.`counts` AS `counts`,`b`.`AVG(avgPrice)` AS `AVG(avgPrice)`,`b`.`AVG(price)` AS `AVG(price)`,`b`.`address1` AS `address1`,(`a`.`avgsalarys` / `b`.`AVG(avgPrice)`) AS `gfyl` from (`近12月薪资` `a` left join `近12月城市房价` `b` on((`a`.`city` = `b`.`address1`))) order by (`a`.`avgsalarys` / `b`.`AVG(avgPrice)`) desc ;

-- --------------------------------------------------------

--
-- 视图结构 `近12月薪资`
--
DROP TABLE IF EXISTS `近12月薪资`;

CREATE VIEW `近12月薪资`  AS  select avg(`item`.`avgsalary`) AS `avgsalarys`,std(`item`.`avgsalary`) AS `STD(avgsalary)`,`item`.`city` AS `city`,count(1) AS `counts` from `item` where ((`item`.`avgsalary` <> 0) and (`item`.`createdate` between (now() - interval 12 month) and now())) group by `item`.`city` having (`counts` > 300) order by `avgsalarys` desc ;
