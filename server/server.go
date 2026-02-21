package server

import (
	"fmt"
	"log"
	"net"
	"github.com/Dhanushkumar-S-G/ToyDB/toy"
)

func Start(port string, toyDB *toy.Toy) {
	
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	defer listener.Close()

	log.Println("Listening on port: ", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error in accepting the connection: ", err)
			continue
		}

		go handleConnection(conn, toyDB)
	}
}

	
