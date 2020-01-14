## 爬虫概念
访问web服务器，获取指定数据信息
## 工作流程
- 明确目标url
- 发送请求，获取应答数据包
- 保存过滤数据，提取有用信息
- 使用分析得到数据信息

##示例
#### 百度贴吧爬取实例
- 指定用户爬取的起始页 创建working函数
- 使用start end 循环爬取每一页数据
- 获取每一页的url---下一页=前一页+50
- 封装HTTP GET函数，实现httpGet，目的是获取一个url数据内容，通过result返回。
- 创建.html文件，分页保存。

#### 并发版本
- 封装爬取一个网页的内容，到函数中（SpiderPage），修改相关参数
- 在working函数中，for循环启动go程调用，相当于爬取多少个页面，起多少个go子程
- 为防止主go程提前结束，引入chan，实现同步，传入进去spiderPage（chan）
- 在spiderPage结尾处，向channel写内容，channel<-index
- 在working函数中，添加新的for循环，从channel中b不断地读取各个子进程的写入的数据
