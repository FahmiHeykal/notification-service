package websocket

import (
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, userID uint) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }

    client := &Client{
        Hub:    hub,
        Conn:   conn,
        Send:   make(chan []byte, 256),
        UserID: userID,
    }

    client.Hub.Register <- client

    go client.Write()
    go client.Read()
}