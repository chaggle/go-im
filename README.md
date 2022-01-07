---
title: "Go语言实现的轻量级IM项目"
date: 2022-01-04T16:18:59+08:00
tag: ["Golang"]
categories: ["Instant Messaging"]
draft: true


---

# go-IM

## 使用 go 实现的微型服务器

项目地址：[https://github.com/chaggle/go-im](https://github.com/chaggle/go-im)

### V0.1 :建立基础的 main.go server.go，相应功能如下

> main 功能主要为创建服务器以及启动服务器
>
> server 功能有：
>
> 1、创建 server 对象
>
> 2、启动 Server 服务(TCP socket 套接字)
>
> 3、处理链接的业务

### V0.2 :用户上线功能

> user 功能新增
>
> 1、创建 user 对象
>
> 2、监听每个 user 对应的 channel 的消息
>
> server 新增功能
>
> 1、新增 OnlineMap 与 Message 属性
>
> 2、在处理客户端上线的 Handler 创建并添加用户（使用到 OS 中的同步 Lock ）
>
> 3、新增广播消息方法以及监听广播消息的 channel 方法

### V0.3 :用户消息广播机制完善

> server 新增功能
>
> 1、完善 handle 模块处处理业务的方法，启动一个针对与当前客户端的读 goroutine

### V0.4 :用户业务层封装

> 对于用户层业务的层次化、模块化
>
> server 中的 user 业务进行迁移
>
> 1、server 关联
>
> 2、新增 Online、Offline、Domessage 方法

### V0.5 :查询用户名以及用户名修改

> user 新增两个功能，
>
> 1、用户名查询的功能
>
> 2、用户名修改的功能，保证每个用户名唯一

### V0.6 :超时强制下线以及私聊功能

> user 新增两个功能，
>
> 1、设置定时器，超时强制剔除，发消息代表活跃，长时间不发消息代表超时，即强制关闭用户连接
>
> 2、私聊功能，通过获取用户名队列中的用户名，来向用户发起私聊

### V1.0 、v1.1 :客户端基本功能

> 新增 client 客户端，当然并非 GUI 界面版本，也可以使用带 GUI 界面版本，满足通信协议即可
>
> 1、连接建立功能、命令行解析功能
>
> 2、客户端菜单功能的预写


### V1.2 :客户端的相应基本请求

> client 新增
>
> 1、用户名修改请求，通过 io.copy() 阻塞监听的方式进行回显输出
>
> 2、用户进入公聊模式，进行消息的广播与退出公聊模式
> 
>3、用户进行私聊模式，选择用户功能封装以及单独对其发送消息


