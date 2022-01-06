package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //模式选择
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       -1,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error :", err)
		return nil
	}

	client.conn = conn

	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.群聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>请输入合法范围内的数字<<<<")
		return false
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口地址(默认8888)")
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}

		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			//公聊模式
			fmt.Println("已进入公聊模式")
			break

		case 2:
			//私聊模式
			fmt.Println("已进入私聊模式")
			break

		case 3:
			//更新用户名
			fmt.Println("已进入更新用户名功能")
			break
		}
	}
}

func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>连接失败>>>>>>>>")
		return
	}

	fmt.Println(">>>>>>>>连接成功！>>>>>>>>")

	client.Run()
}
