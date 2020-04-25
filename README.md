# 教程

## 安装
* 环境变量
    * Linux & Mac
    ```
    go env -w GOPROXY="https://mirrors.aliyun.com/goproxy/" //设置代理
    export GO111MODULE=on
    ```
    * Windows
    ```
    set GOPROXY="https://mirrors.aliyun.com/goproxy/" // 设置代理
    set GO111MODULE=on
    ```
  
* 安装
```
go mod vendor
```