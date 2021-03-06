/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// where all the magic begins
func main() {

    // read the config
    readConfig()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var conn, _ = upgrader.Upgrade(w, r, nil)
        go handleConnection(conn)
    })

    fmt.Println("Server started on port", port)
    http.ListenAndServe(port, nil)
}

// handles the client connections
func handleConnection(conn *websocket.Conn) {
    var msg interface{}
    for {
        err := conn.ReadJSON(&msg)
        if err != nil {
            conn.Close()
            break
        }

        result := check(msg.(map[string]interface{}))
        conn.WriteJSON(result)
    }
}
