# gin-blog

## 一、依赖管理
```
go get                                                                           //把所有依赖包都安装好
govendor init
govendor add +e
govendor update +external
```

## 二、运行
```
go run main.go
```

## 三、接口测试
```
一、标签接口测试
POST：
http://127.0.0.1:8000/api/v1/tags?name=python&state=1&created_by=test1

GET:
http://127.0.0.1:8000/api/v1/tags

PUT:
http://127.0.0.1:8000/api/v1/tags/2?name=java&state=1&modified_by=yoyo


二、文章接口测试
POST：
http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1

GET：
http://127.0.0.1:8000/api/v1/articles

GET：
http://127.0.0.1:8000/api/v1/articles/1

PUT：
http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0

DELETE：
http://127.0.0.1:8000/api/v1/articles/1


三、验证Token
http://127.0.0.1:8000/auth?username=admin&password=admin

运行结果
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE1MTg3MjAwMzcsImlzcyI6Imdpbi1ibG9nIn0.-kK0V9E06qTHOzupQM_gHXAGDB3EJtJS4H5TTCyWwW8"
  },
  "msg": "ok"
}

四、验证jwt
http://127.0.0.1:8000/api/v1/articles
http://127.0.0.1:8000/api/v1/articles?token=23131

得到token
http://127.0.0.1:8000/auth?username=admin&password=admin

带上token请求文章接口
http://127.0.0.1:8000/api/v1/articles?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJhZG1pbiIsImV4cCI6MTU1MTg4NDI1NywiaXNzIjoiZ2luLWJsb2cifQ.xpWYGDvpmzUFW6u8TVeAjq6qjjjTOjLZ17-S8_uO_m0


```
