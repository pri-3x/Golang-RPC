package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

func main() {
	// Address to this variable will be sent to the RPC server
	// Type of reply should be same as that specified on server
	var reply int64
	args := Args{}

	// DialHTTP connects to an HTTP RPC server at the specified network
	client, err := rpc.DialHTTP("tcp", "192.168.1.129:1234") // error handling
	if err != nil {
		log.Fatal("Client connection error: ", err) //A Fatal Error Log. Describes the fatal error log, its location, and contents.
		//The fatal error log is created when a fatal error occurs. It contains information and the state obtained at the time of the fatal error.
	}

	// Invoke the remote function GiveServerTime attached to TimeServer pointer
	// Sending the arguments and reply variable address to the server as well
	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	// Print the reply from the server
	log.Printf("%d", reply)
}
