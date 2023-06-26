# HotViper

对 viper 进行简单的封装，具体实现功能有：

1. 设置模板
2. 回滚配置文件
3. 热加载
4. 编辑接口
5. 恢复默认配置文件模板

具体使用见`example`文件夹下的实例

```bash
cd example
go mod tidy
go run test.go

```
打开另外一个终端请求
```
$ curl -X GET 'http://localhost:3222/'
# {"code":200,"msg":"success","data":{"captcha":{"img-height":80,"img-width":240,"key-long":6},"casbin":{"model-path":"./resource/rbac_model.conf"},"jwt":{"expires-time":604800,"issuer":"catering","signing-key":"catering"},"mysql":{"config":"charset=utf8mb4\u0026parseTime=True\u0026loc=Local","db-name":"gva","host":"127.0.0.1","log-mode":"","log-zap":false,"max-idle-conns":0,"max-open-conns":0,"password":"root","port":"3306","username":"root"},"redis":{"addr":"192.168.138.128:6379","db":0,"password":""},"system":{"addr":8888,"db-type":"mysql","env":"develop","iplimit-count":15000,"iplimit-time":3600,"oss-type":"local","use-multipoint":false},"zap":{"director":"log","encode-level":"LowercaseColorLevelEncoder","format":"console","level":"info","log-in-console":true,"prefix":"[catering]","show-line":true,"stacktrace-key":"stacktrace"}}}

$ curl -X GET 'http://localhost:3222/edit?value=\{"captcha":\{"img-height":80,"img-width":260,"key-long":7\}\}'
$ curl -X GET 'http://localhost:3222/'

# {"code":200,"msg":"success","data":{"captcha":{"img-height":80,"img-width":260,"key-long":7}}}

$ curl -X GET 'http://localhost:3222/reset'

$ curl -X GET 'http://localhost:3222/rollback'
```