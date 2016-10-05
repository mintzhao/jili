# jili
jili, the best schedule calendar

## swagger

### golang
```
swagger generate server -f ./rest_api.json -t ./golang
```

## UI设计

### 参考APP
* [Coursera](https://appsto.re/cn/zwQ5R.i) Discover Tab的内容,可用于联赛、赛季的展示
* [腾讯体育](https://appsto.re/cn/VVSaI.i) 赛程Tab
* [大麦](https://appsto.re/cn/ANgOE.i) 日历功能

### UI列表
* (暂不需要)三张宣传图,全网赛事赛程搜罗|一键导入系统日历,重要比赛不再错过|绿色无推送
* 一个主界面, 如大麦网日历功能, 一个日历放上边,下边是今天的所有比赛
* 一个右侧边栏,用于展示在主界面的赛事筛选
* 将筛选结果一键导入系统日历

## APP细致规则
* APP暂不用开发管理端,直接数据采集
* APP暂不用设置用户系统
* 收费规则,免前n条导入,从第n+1条起,每条一分/一角钱
* 使用苹果支付 or pingpp(微信、支付宝、Apple Pay)