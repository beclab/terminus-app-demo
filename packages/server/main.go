package main

import (
	"fmt"

	"github.com/beclab/terminus-app-demo/pkg/apiserver"
)

func main() {
	listen := ":9001"
	done := make(chan struct{})

	server := apiserver.New(listen)

	go server.Start()

	fmt.Printf("Server started on %s\n", listen)

	<-done

	server.Stop()
}
