package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		// 1, 2, ...99
		//1-> Open or closed
		//Solo sitios de pruebas, puede ser considerado un ataque.
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}

		conn.Close()
		fmt.Printf("Port %d is open", i)
		fmt.Println("")
	}
}
