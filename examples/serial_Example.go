package main

import (
"fmt"
"gopkg.in/vmihailenco/msgpack.v2"
)

func main() {
    ExampleMarshal()
    
}
func ExampleMarshal() {
    b, err := msgpack.Marshal("a|b")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", b)
    // Output:

    var out bool
    err = msgpack.Unmarshal([]byte{0xc3}, &out)
    if err != nil {
        panic(err)
    }
    fmt.Println(out)
    // Output: []byte{0xc3}
    // true
}
