/*
	This program is the main logic of my server.
	This server listen to requests and handles
	all of request.
	Also I am using concurrency to improve
	my server efficiency!

	Coded by: Pouya Mohammadi
	github: pouyam79i
*/

/*
TODO: remove this comment!!!!!!! *******************************

	this code is a review on how to create http server using golang!
	The article below gives you the base idea!
	https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

*/

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

// this method is for testing our http res to get req!
func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Hello Req from: ", ctx.Value(keyServerAddr))

	var res_msg string = "Hello "

	// TODO: uncomment if you want to read query for name value
	// Check if query exists!
	// if r.URL.Query().Has("name") {
	// 	var name string = r.URL.Query().Get("name")
	// 	if name == "" {
	// 		res_msg += "World"
	// 	} else {
	// 		res_msg += name
	// 	}
	// } else {
	// 	res_msg += "World"
	// }

	// Check the req body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Could not read body")
	} else {
		fmt.Println("Received body:\n", body)
	}

	// TODO: uncomment if you want to use form to get name value
	// check values from post (maybe user was filing a form)
	name := r.PostFormValue("name")
	// TODO: if you want to retrieve value from form or query uncomment line below and comment line above.
	// name = r.FormValue("name")
	if name == "" {
		// you can add costume header!
		w.Header().Set("x-missing-field", "name")
		// You can set and http status code!
		w.WriteHeader(http.StatusBadRequest)
		return
		// name = "World"
	}

	io.WriteString(w, res_msg+name+"!\n")
}

// TODO: return the main html file
// this method returns homepage of our website!
func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Root Req from: ", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is from ROOT '/'\n")
}

// TODO: complete this
// http req listener!
func ListenAndServe(ip, port string, handler http.Handler) error {
	err := http.ListenAndServe(ip+":"+port, handler)
	return err
}

func main() {
	fmt.Println("Server is booting...")

	// attaching handler functions
	mux := http.NewServeMux() // server mux instead of default http handler
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	// TODO: uncomment if you want only one http server running here
	// // running single http server
	// fmt.Println("Server is on!")
	// err := ListenAndServe("", "8081", mux)
	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Println("Server is shut down")
	// } else if err != nil {
	// 	// Print unexpected errors!
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// TODO: This is a multi server program - remove comment of Server two if needed
	// here we have our multi server logic:
	ctx := context.Background()
	// TODO: if you want more than one server uncomment line below and comment above line.
	// ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":8081",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	// serverTwo := &http.Server{
	// 	Addr:    ":8082",
	// 	Handler: mux,
	// 	BaseContext: func(l net.Listener) context.Context {
	// 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
	// 		return ctx
	// 	},
	// }

	// go func() {
	fmt.Println("Server One is On!")
	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server One is down!")
	} else if err != nil {
		fmt.Println("Error on server one: ", err)
	}
	// cancelCtx()
	// }()

	// go func() {
	// 	fmt.Println("Server Two is On!")
	// 	err := serverTwo.ListenAndServe()
	// 	if errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Println("Server Two is down!")
	// 	} else if err != nil {
	// 		fmt.Println("Error on server two: ", err)
	// 	}
	// 	cancelCtx()
	// }()
	// <-ctx.Done()
	fmt.Println("Host is down!")

}
