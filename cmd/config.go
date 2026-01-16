package cmd

const (
	ServerIP   = "127.0.0.1"
	ServerPort = "9999"
)

func GetServerAddress() string {
	return ServerIP + ":" + ServerPort
}
