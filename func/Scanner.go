package _func

import (
	"fmt"
	"net"
)

// Scanning function
func ScanUnique(port string, ip string, network string, nbr_open int) int {
	address := ip + ":" + port

	connection, err := net.Dial(network, address)

	if err != nil {
		fmt.Println("[-] Port " + port + " is closed. [-]")
	} else {
		fmt.Println("[+] Connection established : " + connection.RemoteAddr().String())
		nbr_open++
	}
	return nbr_open
}
