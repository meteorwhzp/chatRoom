package main

import (
	logger "github.com/shengkehua/xlog4go"
	"net"
	"bufio"
	"runtime"
	"time"
	"flag"
	"math/rand"
	"fmt"
)


type client chan<-string //对外发送消息的通道

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
	logconf = flag.String("l", "./conf/log.json", "log config file path")
)



func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	// log
	if err := logger.SetupLogWithConf(*logconf); err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("Starting service ...")

	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		logger.Error(err.Error())
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		go handleConn(conn)
	}

}

func broadcaster() {
	clients := make(map[client]bool) //所有连接的客户端
	for {
		select {
		case msg := <- messages:
			//把所有接收的消息广播给所有的客户端
			//发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "you are" + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		logger.Info("conn %v : %s", conn, msg)
		fmt.Fprintln(conn, msg)
	}
}
