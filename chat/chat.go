package chat

import (
	"log"

	"golang.org/x/net/context"

)

type Server struct {

}

func (s *Server ) SayHello(ctx context.Context, message *Message)(*Message, error){
	log.Printf("Mensaje recibido del cliente: %s", message.Body)
	return &Message{Body: "Hola desde el servidor"}, nil
}

