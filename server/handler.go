package server

import (
	"bufio"
	"github.com/Dhanushkumar-S-G/ToyDB/command"
	"github.com/Dhanushkumar-S-G/ToyDB/toy"
	"net"
	"time"
)

func handleConnection(conn net.Conn, toyDB *toy.Toy) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5*time.Minute))

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		input := scanner.Text()
		
		cmd, args, err := command.ParseCommand(input)
		if err != nil {
			conn.Write([]byte("Error: " + err.Error() + "\n"))
			continue
		}
		
		result, err := command.ExecuteCommand(cmd, args, toyDB)
		if err != nil {
			conn.Write([]byte("Error: " + err.Error() + "\n"))
			continue
		}
		
		conn.Write([]byte(result + "\n"))
	}
}