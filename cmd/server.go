package cmd

import (
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var logServer *os.File

const (
	// 清理超时客户端，监听心跳时间
	t1 = 1 * time.Second

	// 心跳超时时间
	t2 = 10 * time.Second
)

func loggerServer() {
	var err error
	logServer, err = os.OpenFile("heartbeat_server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logServer)
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Heart beat server",
	Run: func(cmd *cobra.Command, args []string) {
		loggerServer()
		receiveHeartBeat()
	},
}

func receiveHeartBeat() {
	addr, _ := net.ResolveUDPAddr("udp", GetServerAddress())
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	var clients sync.Map
	buffer := make([]byte, 65507)

	log.Println("心跳监听开始")

	go func() {
		ticker := time.NewTicker(t1)
		defer ticker.Stop()

		for range ticker.C {
			clients.Range(func(key, value interface{}) bool {
				if time.Since(value.(time.Time)) > t2 {
					clients.Delete(key)

					elapsed := time.Since(value.(time.Time))
					log.Printf("%s 心跳超时：%v秒，阈值：%v秒\n", key, elapsed, t2)

					block()
					sendCanMsg()
				}
				return true
			})
		}
	}()

	for {
		n, clientAddr, _ := conn.ReadFromUDP(buffer)
		if n > 0 && string(buffer[:n]) == "HEARTBEAT" {
			clientIP := clientAddr.String()

			_, loaded := clients.LoadOrStore(clientIP, time.Now())

			if !loaded {
				log.Printf("%s 心跳第一次收到\n", clientIP)
			} else {
				clients.Store(clientIP, time.Now())
			}
		}
	}
}

func block() {
	if logServer != nil {
		logServer.Sync()
	}
}

func sendCanMsg() {
	cmd := exec.Command("sh", "-c", "cansend can0 011#0000040000000040")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	go func() {
		err := cmd.Wait()
		if err != nil {
			panic(err)
		}
	}()
}
