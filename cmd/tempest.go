package main

import (
	"fmt"
	"go-tempest/tempest"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	network := tempest.NewNetwork("tempest")
	if err := network.Start(net.IPv4zero); err != nil {
		panic(err)
	}

	defer func() {
		fmt.Printf("Stopping %s network\n", network.NetworkName)
		if err := network.Stop(); err != nil {
			panic(err)
		}
	}()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	<-sigChannel
	fmt.Println("\nShutting down...")
}
