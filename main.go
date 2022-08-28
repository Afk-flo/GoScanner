package main

import (
	Scanner "TCP_Scanner/func"
	"flag"
	"fmt"
	"strconv"
	"sync"
)

func main() {

	// Ask for an IP or a specific port or even a specific network protocol
	ip := flag.String("h", "", "Select an IP address to scan")
	uniquePort := flag.String("p", "-1", "Select a unique port to scan")
	network := flag.String("n", "tcp", "Select a protocol")
	full := flag.String("f", "none", "Select if you want a run on the 65000 ports, otherwise it'll run on the 1024 first ports")
	flag.Parse()

	minPorts := 1024
	maxPorts := 65000

	// Generate port

	// Initialisation
	fmt.Println(Scanner.Purple + "[*] Scan initialisation [*]" + Scanner.Reset)

	// handle thread
	var wg sync.WaitGroup

	// Simple port
	if *uniquePort != "-1" {
		wg.Add(1)
		go Scanner.ScanUnique(*uniquePort, *ip, *network, &wg)

		// TOP 100 ports
	} else if *full == "none" {
		for i := 0; i < minPorts; i++ {
			wg.Add(1)
			go Scanner.ScanUnique(strconv.Itoa(i), *ip, *network, &wg)
		}

		// 65000 ports
	} else {
		// Don't use that before I put the thread in place ..
		for i := 0; i < maxPorts; i++ {
			wg.Add(1)
			go Scanner.ScanUnique(strconv.Itoa(i), *ip, *network, &wg)
		}
	}

	wg.Wait()

	// End of the script
	fmt.Println(Scanner.Purple + "[*] Scan is finished using the " + *network + " protocol on the " + *ip + " [*]" + Scanner.Reset)
	fmt.Println(" ")

}
