---
title: "Go语言实现的轻量级IM项目"
date: 2022-01-04T16:18:59+08:00
tag: ["Golang"]
categories: ["Instant Messaging"]
draft: true
---

# go-IM

## 使用 go 实现的微型服务器

项目地址：https://github.com/chaggle/go-im

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
> 1、完善 handle 模块处处理业务的方法，启动一个针对与当前客户端的读goroutine
