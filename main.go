package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var port =  3000

func main() {
    initializeStore()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var conn, _ = upgrader.Upgrade(w, r, nil)
        var msg interface{}
        go func(conn *websocket.Conn) {
            for {

                err := conn.ReadJSON(&msg)
                if err != nil {
                    conn.Close()
                    break
                }

                result := check(msg.(map[string]interface{}))
                conn.WriteJSON(result)
            }
        }(conn)
    })

    fmt.Println("Server started on port", port)
    http.ListenAndServe(":3000", nil)
}
