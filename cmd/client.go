package cmd

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var logClient *os.File

func loggerClient() {
	var err error
	logClient, err = os.OpenFile("heartbeat_client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logClient)
}

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Heart beat client",
	Run: func(cmd *cobra.Command, args []string) {
		loggerClient()
		sendHeartBeat()
	},
}

func sendHeartBeat() {
	serverAddr, _ := net.ResolveUDPAddr("udp", GetServerAddress())
	localAddr, _ := net.ResolveUDPAddr("udp", GetClientAddress())

	conn, _ := net.DialUDP("udp", localAddr, serverAddr)
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	heartbeatMsg := []byte("HEARTBEAT")

	log.Println("心跳开始，端口： " + GetClientPort())

	for range ticker.C {
		_, err := conn.Write(heartbeatMsg)
		if err != nil {
			continue
		}
	}
}
