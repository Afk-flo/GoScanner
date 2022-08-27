package main

import (
	Scanner "TCP_Scanner/func"
	"flag"
	"fmt"
	"strconv"
)

// Var prep
var nbr_open int = 0

func main() {

	// Ask for an IP or a specific port or even a specific network protocol
	ip := flag.String("h", "", "Select an IP address to scan")
	unique_port := flag.String("p", "-1", "Select a unique port to scan")
	network := flag.String("n", "tcp", "Select a protocol")
	full := flag.String("f", "none", "Select if you want a run on the 65000 ports, otherwise it'll run on the top 30 ports")
	flag.Parse()

	// ports definition
	port := [30]string{
		"20", "49", "63",
		"21", "22", "23",
		"25", "53", "79",
		"80", "88", "110",
		"111", "135", "137",
		"443", "995", "993",
		"631", "465", "143",
		"123", "14147", "105",
		"66", "143", "2224",
		"8000", "8888", "3306"}

	maxPorts := 65000

	// Simple bug slow alg
	if *unique_port != "-1" {
		nbr_open = Scanner.ScanUnique(*unique_port, *ip, *network, nbr_open)
	} else if *full == "none" {
		for _, p := range port {
			nbr_open = Scanner.ScanUnique(p, *ip, *network, nbr_open)
		}
	} else {
		// Don't use that before I put the thread in place ..
		for i := 0; i < maxPorts; i++ {
			nbr_open = Scanner.ScanUnique(string(rune(i)), *ip, *network, nbr_open)
		}
	}

	// End of the script
	fmt.Println(" ")
	fmt.Println("[*] Scan is finished with " + strconv.Itoa(nbr_open) + " ports open using the " + *network + " protocol [*]")
	fmt.Println(" ")

}
