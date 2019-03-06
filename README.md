# gin-blog

## 依赖管理
```
govendor init
govendor add +e
govendor update +external
```

## 运行
```
go run main.go
```

## 接口测试
```
post测试
http://127.0.0.1:8000/api/v1/tags?name=python&state=1&created_by=test1

get测试
http://127.0.0.1:8000/api/v1/tags

put测试
http://127.0.0.1:8000/api/v1/tags/2?name=java&state=1&modified_by=yoyo
```
