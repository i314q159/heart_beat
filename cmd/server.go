package cmd

import (
	"fmt"
	"net"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Heart beat server",
	Run: func(cmd *cobra.Command, args []string) {
		receive_heart_beat()
	},
}

func receive_heart_beat() {
	addr, _ := net.ResolveUDPAddr("udp", ":"+ServerPort)
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	var clients sync.Map
	buffer := make([]byte, 1024)

	fmt.Println("Heartbeat receiver started on port " + ServerPort)

	// 清理超时客户端，每1秒检查一次，超时时间是5秒
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			clients.Range(func(key, value interface{}) bool {
				if time.Since(value.(time.Time)) > 5*time.Second {
					clients.Delete(key)
					fmt.Printf("Client %s Timeout and Removed\n", key)
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
				fmt.Printf("First Connection From %s\n", clientIP)
			} else {
				clients.Store(clientIP, time.Now())
				// fmt.Printf("HEARTBEAT RECEIVED from %s\n", clientIP)
			}
		}
	}
}

func sendCanMsg() {
	cmd := exec.Command("bash", "-c", "cansend can0 011#0000040000000040; echo 'CAN指令已发送';")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	fmt.Println("CAN 指令已发送")
}
