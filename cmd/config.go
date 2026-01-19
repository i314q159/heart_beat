package cmd

const (
	ServerIP   = "192.168.10.2"
	ServerPort = "9999"
)

func GetServerAddress() string {
	return ServerIP + ":" + ServerPort
}
