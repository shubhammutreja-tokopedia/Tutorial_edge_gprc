package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("Recived mssg : %s", msg.Body)
	return &Message{Body: "Hello from server go"}, nil
}
