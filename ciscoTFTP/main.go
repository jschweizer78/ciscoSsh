package main

import (
	"fmt"
	"os"
)

// TODO fix menu for background. Then if background what type

func main() {
	var subFolder string
	var serverIP string

	var answerRunType int

	fmt.Println("What would you like the subfolder name to be?")
	fmt.Scan(&subFolder)

	fmt.Println(`What is the server hostname/IP?`)
	fmt.Scan(&serverIP)

	fmt.Println(`What would you like to do?
Please choose one of the following:
1 Standard Ring list
2 Distinctive Ring
3 Image download`)
	fmt.Scan(&answerRunType)

	switch answerRunType {
	case 1, 2:
		getStandardRings(subFolder, serverIP, answerRunType)
	case 3:
		fmt.Println("Under construction")
	case 99:
		fmt.Println("Thanks for playing")
		os.Exit(1)
	default:
		fmt.Println("No valid menu option selected")
		os.Exit(1)
	}
}

/*
c, err := tftp.NewClient("172.16.4.21:69")
wt, err := c.Receive("foobar.txt", "octet")
file, err := os.Create(path)
// Optionally obtain transfer size before actual data.
if n, ok := wt.(IncomingTransfer).Size(); ok {
    fmt.Printf("Transfer size: %d\n", n)
}
n, err := wt.WriteTo(file)
fmt.Printf("%d bytes received\n", n)
*/
