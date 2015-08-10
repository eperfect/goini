# go-config
基于Golang的轻量级ini配置文件解析库。
A light weight library parsing config.ini file for Golang

##安装：

```
go get github.com/eperfect/goini/
```

###INI配置文件格式：
```
ip=localhost

[server]
#server config
ip=192.168.1.1
port=8080


[database]
 ip=		localhost
port= 6379
```
###说明：
1.＝配置属性
2.[]配置分区
3.#后为文件注释
4.配置属性中忽略空格和Tab

###代码中使用：
```
InitConfig("config.ini")  //初始化配置文件，参数为配置文件路径

GetValue("", "ip")        //获取配置信息，分区（Section）不填写的话默认为”default“

GetValue("server", "ip")  //获取指定分区的字段配置信息

SetValue("server", "port", "9090")  //设置分区字段属性

GetValue("invalid_section", "ip")   //获取不存在的配置（分区或字段不存在）返回空字符串
```

