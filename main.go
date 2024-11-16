package main

import (
	"os"
	"strconv"
	"clickdata/server"
)

const (
	defaultAddr = "127.0.0.1"
	defaultPort = "8000"
	addrEnv     = "CLICK_ADDR"
	portEnv     = "CLICK_PORT"
)

var serverAddr string

func main() {
	address := func(defaultAddr string) string {
		result := os.Getenv(addrEnv)
		if result == "" {
			result = defaultAddr
		}
		return result
	}(defaultAddr)

	port, err := func(defaultPort string) (int, error) {
		result := os.Getenv(portEnv)
		if result == "" {
			result = defaultPort
		}
		return strconv.Atoi(result)
	}(defaultPort)

	if err != nil {
		panic(err)
	}

	server.NewServer(address, port).Run()
}
