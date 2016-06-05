package main

import (
"net"
"fmt"
"log"
"gopkg.in/vmihailenco/msgpack.v2"
)

func main() {
	fmt.Println("Server is running on 1500");
	listener,err := net.Listen("tcp","localhost:1500")
	if err!= nil{
		log.Fatalln(err)
	}

	defer listener.Close() // close the listener later

	for {
		conn,err := listener.Accept()
		if err!= nil{
			log.Fatalln(err)
		}
		fmt.Println("Started");
		go listenConnection(conn)	// spawn os threads
	}
	
}

// Each connection will have its own thread
func listenConnection(conn net.Conn) {
	for {
		var out string
		buffer := make([]byte,1400)
		dataSize,err := conn.Read(buffer)
		if err!= nil{
			log.Fatalln("Connection closed")
		}
		data := buffer[:dataSize]
    	err = msgpack.Unmarshal(data, &out)
    	if err != nil {
        	panic(err)
    	}
		fmt.Println("received a message",string(out))
		_,err = conn.Write(data)

		if err!= nil{
			log.Fatalln(err)

		} 
		fmt.Println("message sent",string(out))
	}
}
