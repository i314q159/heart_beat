package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Heart beat client",
	Run: func(cmd *cobra.Command, args []string) {
		send_heart_beat()
	},
}

func send_heart_beat() {
	serverAddr, _ := net.ResolveUDPAddr("udp", GetServerAddress())
	conn, _ := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	heartbeatMsg := []byte("HEARTBEAT")

	fmt.Println("Heartbeat Client Started")

	for range ticker.C {
		_, err := conn.Write(heartbeatMsg)
		if err != nil {
			fmt.Printf("Failed to send heartbeat: %v\n", err)
			continue
		}

		// fmt.Println("HEARTBEAT SENT")
	}
}
