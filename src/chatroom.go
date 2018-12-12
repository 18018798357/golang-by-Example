// 聊天室
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// 广播
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 向所有人广播收到的信息
			// 客户的外发消息渠道
			for cli := range clients {
				cli <- msg
			}
		// 到来
		case cli := <-entering:
			clients[cli] = true
		// 离开
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// 处理连接
func handleConn(conn net.Conn) {
	// 传出客户端消息
	ch := make(chan string)
	go clientWriter(conn, ch)
	// 获取ip
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "has arrived"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		// 注意:忽略网络错误
		fmt.Fprintf(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
