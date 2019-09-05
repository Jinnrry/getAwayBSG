// 这里是分析数据时用的的查询语句


// 房价均价查询语句
db.lianjia.aggregate([
	{
		'$match': {
			$and: [
				{ "address.0": { $exists: true } },
				//统计数据的的时间区间，使用的挂牌时间统计
				{
					"guapaitime": {
						$gt: ISODate("2019-07-01T00:00:00.000+0000"),
						$lt: ISODate("2019-08-01T00:00:00.000+0000")
					}
				},
				{ "address.0": "成都" }
			]
		},
	},
	{
		$group: {
			_id: { $substr: ["$Link", 0, 22] },
			city: { $first: "$address" },
			count: { $sum: 1 },
			avg_UnitPrice: { $avg: "$UnitPrice" },
			std: { $stdDevPop: "$UnitPrice" },
		}
	},
	{
		$project:
			{
				count: 1,        //总数
				avg_UnitPrice: 1, //每平米均价
				std: 1,   //标准差
				ratio: { $divide: ["$std", "$avg_UnitPrice"] }, //标准差与均价的比值
				city: { $slice: ['$city', 0, 1] }
			}
	},
	{
		'$sort': { count: -1 }
	}
]);


// 平均薪资查询语句
db.zhilian.aggregate([
	{'$match': {"workingExp.name": "1-3年"}},
	{
		$group: {
			_id: {"$arrayElemAt": ["$city.items.name", 0]},
			count: {$sum: 1},
			avg: {$avg: "$avg"},
			std: {$stdDevPop: "$avg"},
		}
	},
	{
		$project:
			{
				count: 1,   //总数
				avg: 1, //平均薪资
				std: 1,   //标准差
				ratio: {$divide: ["$std", "$avg"]} //标准差与均价的比值
			}
	},
	{
		'$sort': {count: -1}
	}
]);


// 薪资，按公司规模和城市分组
db.zhilian.aggregate([
	{
		$group: {
			_id: {city: {"$arrayElemAt": ["$city.items.name", 0]}, "公司规模": "$company.size.name"},
			count: {$sum: 1},
			avg: {$avg: "$avg"},
			std: {$stdDevPop: "$avg"},
		}
	},
	{
		$project:
			{
				count: 1,   //总数
				avg: 1, //平均薪资
				std: 1,   //标准差
				ratio: {$divide: ["$std", "$avg"]} //标准差与均价的比值
			}
	},
	{
		'$sort': {count: -1}
	}
]);


// 当前进度
db.lianjia.aggregate([
	{
		'$sort': {detailCrawlTime: -1}
	}
], {allowDiskUse: true});


// 租房数据
db.lianjia_zufang.aggregate([
	{'$match': {"mianji": {$gt: 0}}},
	{
		$group: {
			_id: "$city",
			count: {$sum: 1},
			avg: {$avg: "$price"},
			std: {$stdDevPop: "$price"},
			unitPrice: {$avg: {$divide: ["$price", "$mianji"]}}
		}
	},
	{
		$project: {
			unitPrice: 1,  // 单位价格
			count: 1,        //总数
			avg: 1, //每平米均价
			std: 1,   //标准差
			ratio: {$divide: ["$std", "$avg"]} //标准差与均价的比值
		}
	},
	{
		'$sort': {count: -1}
	}
]);

