package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//online user list
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//message channel
	Message chan string
}

//create a server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部的在线User
		this.mapLock.Lock() //OS同步的机制
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCat(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handle(conn net.Conn) {
	/* fmt.Println("Connecting Success") */

	user := NewUser(conn, this)

	//用户上线
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//建立一个func函数来接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//用户下线
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//提取用户的消息(去除\n)
			msg := string(buf[:n-1])

			//广播得到的相关消息！
			user.Domessage(msg)

			//只要用户发消息，就赋予活跃状态
			isLive <- true

		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//当前用户活跃，重置定时器,不做任何事情，为了激活select，更新定时器
		case <-time.After(time.Second * 300):
			//超时，强制关闭当前的channel
			user.SendMsg("用户下线，请重新登陆")

			//销毁资源
			close(user.C)

			//关闭连接
			conn.Close()

			//退出当前Handler
			return //runtime.Goexit()
		}
	}

}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err : ", err)
		return
	}

	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err : ", err)
			continue
		}

		//do handle
		go this.Handle(conn)
	}
}
