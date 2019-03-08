# gin-blog项目使用文档

## 一、项目初始化
初始工作区

首先，我们需要增加一个工作区（GOPATH）路径用于我们的gin-blog项目。

将你新的工作区加入到/etc/profile中的GOPATH环境变量中， 并在新工作区中，建立bin、pkg、src三个目录。

```
mkdir -p /opt/path/
export GOPATH=/opt/path/
mkdir -p /opt/path/{bin,pkg,src}

echo "export GOPATH=/opt/path/" >>/etc/profile
```
## 二、克隆项目
```
cd /opt/path/src

git clone https://github.com/Lancger/gin-blog.git
```

## 三、依赖管理
```
#下载依赖包管理工具
go get -u -v github.com/kardianos/govendor

#将包管理工具同步到$GOROOT的bin目录，或者将$GOPATH的BIN目录给加入环境变量中
cp /opt/path/bin/govendor /usr/local/go/bin/

#检验govendor版本
govendor -version
$ v1.0.9

go get                 //把所有依赖包都安装好
govendor init          //会生成一个vendor目录和vendor.json, 里面并没有依赖包信息
govendor add +external
govendor add +e       //#将GOPATH中本工程使用到的依赖包自动移动到vendor目录中,如果本地GOPATH没有依赖包，先go get相应的依赖包

govendor update +external
```

## 四、运行
```
go run main.go
```

## 五、接口测试
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
