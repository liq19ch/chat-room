package main

type room struct {
	// holds incoming msg
	forward chan []byte
	join    chan *client
	leave   chan *client
	// all current clinets in this room
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		// if there's a client wants to join, save it to map
		case client := <-r.join:
			r.clients[client] = true
		// if wants to move, delete the client from map & close the client's send channel
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		// if there's msg wants to send in to this room, add them to send channels of all client
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}
