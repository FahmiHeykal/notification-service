package websocket

type Hub struct {
    Clients    map[*Client]bool
    Register   chan *Client
    Unregister chan *Client
    Broadcast  chan []byte
}

func NewHub() *Hub {
    return &Hub{
        Clients:    make(map[*Client]bool),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Broadcast:  make(chan []byte),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.Clients[client] = true
        case client := <-h.Unregister:
            if _, ok := h.Clients[client]; ok {
                delete(h.Clients, client)
                close(client.Send)
            }
        case message := <-h.Broadcast:
            for client := range h.Clients {
                select {
                case client.Send <- message:
                default:
                    close(client.Send)
                    delete(h.Clients, client)
                }
            }
        }
    }
}

func (h *Hub) SendToUser(userID uint, message string) {
    for client := range h.Clients {
        if client.UserID == userID {
            client.Send <- []byte(message)
        }
    }
}