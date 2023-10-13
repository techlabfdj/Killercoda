package main

import (
	myapi "gitlab-techlab/techlab/training/golang/gin-samples/datas/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const listenAddr = "127.0.0.1:8080"
const logDir ="logs"

func main() {

	apiServer := new(myapi.Server)
	if err := apiServer.Init(listenAddr, logDir); err != nil {
		log.Println("api server initialization failed:", err)
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		apiServer.Startup()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down Server ...")
	apiServer.Shutdown(5)
	log.Println("Server shutdown completed")

}
