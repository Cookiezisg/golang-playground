package main

import (
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var reply string
	err = client.Call("PersonService.Greet", "Hello", &reply)
	if err != nil {
		panic(err)
	}

	println("Response from server:", reply)
}
