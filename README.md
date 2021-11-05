
	 3970  go env -w GOPROXY=https://goproxy.cn,direct

	  3976  mkdir SMZDTZ
 3977  cd SMZDTZ
 3978  ls
 3979  mkdir go
 3980  cd go
 3981  pwd
 3982  go env -w GOPATH=/Users/aha/SMZDTZ/go

  3991  go env -w GO111MODULE=on
 3992  go env -w GOPROXY=https://goproxy.cn,direct
 3993  go mod init
	
	npm i nodemon -g

	nodemon --exec go run hello.go --signal SIGTERM