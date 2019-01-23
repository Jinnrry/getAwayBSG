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
 `ljID` double NOT NULL COMMENT '链家ID',
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