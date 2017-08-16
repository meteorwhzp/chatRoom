package main

import (
	"net"
	logger "github.com/shengkehua/xlog4go"
	"io"
	"os"
	"flag"
	"bufio"
)
var (
	logconf = flag.String("l", "./conf/log.json", "log config file path")

)

func main() {
	// log
	if err := logger.SetupLogWithConf(*logconf); err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("Starting service ...")

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		logger.Error(err.Error())
	}
	//done := make(chan struct{})
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			logger.Info(input.Text())
		}
	}()

	go func() {
		mustCopy(conn, os.Stdin)
	}()
	select {

	}
	conn.Close()

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		logger.Error(err.Error())
	}
}
