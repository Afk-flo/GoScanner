package _func

import (
	"fmt"
	"net"
	"sync"
)

var Red = "\033[31m"
var Green = "\033[32m"
var Reset = "\033[0m"
var Purple = "\033[35m"

// Scanning function
func ScanUnique(port string, ip string, network string, wg *sync.WaitGroup) {
	defer wg.Done()
	address := ip + ":" + port
	connection, err := net.Dial(network, address)

	// handle closed port
	if err != nil {
		return
	}

	fmt.Println(Green + "[+] Port : " + port + " is open. [+]" + Reset)

	// handle closing
	err2 := connection.Close()
	if err2 != nil {
		return
	}
}
