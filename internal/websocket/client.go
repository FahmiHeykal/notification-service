package websocket

import (
    "github.com/gorilla/websocket"
)

type Client struct {
    Hub      *Hub
    Conn     *websocket.Conn
    Send     chan []byte
    UserID   uint
}

func (c *Client) Read() {
    defer func() {
        c.Hub.Unregister <- c
        c.Conn.Close()
    }()

    for {
        _, _, err := c.Conn.ReadMessage()
        if err != nil {
            break
        }
    }
}

func (c *Client) Write() {
    defer func() {
        c.Conn.Close()
    }()

    for message := range c.Send {
        err := c.Conn.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            break
        }
    }
}