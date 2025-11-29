package main

import (
	"net"
	"net/rpc"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet(greeting string, reply *string) error {
	*reply = greeting + ", " + p.Name + "! Currently you are " + strconv.Itoa(p.Age) + " years old."
	p.Age += 1 // 每次打招呼时年龄加一
	return nil
}

func main() {
	// 注册 RPC 服务

	newPerson := new(Person)
	newPerson.Name = "Weilin"
	newPerson.Age = 30

	err := rpc.RegisterName("PersonService", newPerson)
	if err != nil {
		panic(err)
	} else {
		println("RPC service registered successfully")
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	println("RPC server listening on :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
