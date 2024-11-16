package main

import (
	c "clickdata/components"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

const (
	defaultAddr = "127.0.0.1"
	defaultPort = "8000"
	addrEnv     = "CLICK_ADDR"
	portEnv     = "CLICK_PORT"
)

var serverAddr string

func init() {
	serverAddr = func(defaultAddr, defaultPort string) string {
		addr := os.Getenv(addrEnv)
		if addr == "" {
			addr = defaultAddr
		}
		port := os.Getenv(portEnv)
		if port == "" {
			port = defaultPort
		}
		return fmt.Sprintf("%s:%s", addr, port)
	}(defaultAddr, defaultPort)

}

func main() {
	count := 0

	http.Handle("/", templ.Handler(c.Index()))

	http.HandleFunc("POST /click", func(w http.ResponseWriter, r *http.Request) {
		count += 1
		w.Write([]byte(fmt.Sprintf("%d", count)))
	})

	log.Printf("Serving on %s", serverAddr)

	http.ListenAndServe(serverAddr, nil)
}
