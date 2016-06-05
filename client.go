package main

import (
"net"
"fmt"
"log"
"gopkg.in/vmihailenco/msgpack.v2"
)

func main() {
    conn,err := net.Dial("tcp","localhost:1500")
    if err!= nil{
        log.Fatalln(err)
    }

    message, err := msgpack.Marshal("a|b")
    if err != nil {
        panic(err)
    }

    _,err = conn.Write(message)
    if err!= nil{
	log.Fatalln(err)
    }

    fmt.Println("Message sent")
    for {
        var out string
        buffer := make([]byte,1400)
        dataSize,err := conn.Read(buffer)	
        if err!= nil{
            log.Fatalln("Connection closed")
            return
        }

        data := buffer[:dataSize]
        err = msgpack.Unmarshal(data, &out)
        if err != nil {
            panic(err)
        }
	    fmt.Println("received a message",string(out))
	
        }

}