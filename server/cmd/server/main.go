/*
	This program is the main logic of my server.
	This server listen to requests and handles
	all of request.
	Also I am using concurrency to improve
	my server efficiency!

	Coded by: Pouya Mohammadi
	github: pouyam79i
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

const keyServerAddr = "web-service.ir"

// server builder
func BuildServer(ip, port string) {
	fmt.Println("Server is booting...")
	// attaching handler functions
	mux := http.NewServeMux() // server mux instead of default http handler
	mux.Handle("/", http.FileServer(http.Dir("../../web/static")))

	// setting server config
	ctx := context.Background()
	addr := ip + ":" + port
	serverOne := &http.Server{
		Addr:    addr,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	// running server
	fmt.Println("Server One is On!")
	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server One is down!")
		err = nil
	} else if err != nil {
		fmt.Println("Error on server one: ", err)
	}

}

// Running application
func main() {
	BuildServer("", "8080")
	fmt.Println("Host is down!")
}
