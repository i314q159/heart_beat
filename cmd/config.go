package cmd

const (
	serverIP   = "192.168.10.2"
	serverPort = "9999"
	clientPort = "8888"
)

func GetServerAddress() string {
	return serverIP + ":" + serverPort
}

func GetClientPort() string {
	return clientPort
}

func GetServerPort() string {
	return serverPort
}
