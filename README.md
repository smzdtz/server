### 单元测试
```bash
# 在package下执行
$ cd ./datacenter/eastmoney

# -v 参数显示每个用例的测试结果，-cover 参数可以查看覆盖率
$ go test

# 执行某个测试用例
$ go test -run TestQueryFundInfo
```

### 其他
1. vscode自动重启
```bash
$ npm i nodemon -g
$ nodemon --exec go run hello.go --signal SIGTERM
```

Paste JSON as Code