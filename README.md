## 1. 修改配置

`config/config.yaml`

确保代理可用

## 2. 编译运行

```golang
go mod tidy

go build -o main cmd/main.go

main
```

## 3. 统计问题数

```shell
cd pkg/utils

go test
``` 
