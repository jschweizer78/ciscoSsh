package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pin/tftp"
)

func readRingListFile(path string) CiscoIPPhoneRingList {
	var ringType CiscoIPPhoneRingList

	fileSlice, err := ioutil.ReadFile(path)
	check(err)

	err = xml.Unmarshal(fileSlice, &ringType)
	check(err)

	return ringType
}

func getTftpFile(fileName, subFolder, serverIP string, result chan string) {
	serverString := fmt.Sprintf("%s:69", serverIP)

	c, err := tftp.NewClient(serverString)
	check(err)

	// fmt.Println("connected to ", serverString)
	wt, err := c.Receive(fileName, "octet")
	check(err)

	// fmt.Println("file recived ")
	path := filepath.Join(subFolder, fileName)

	currentPath := "./" + path
	// fmt.Println(currentPath)
	file, err := os.Create(currentPath)
	check(err)

	n, err := wt.WriteTo(file)
	check(err)

	// fmt.Printf("%d bytes received\n", n)
	log := fmt.Sprintf("Connected to %s and %s received. It containted %d bytes, and was saved to %s", serverString, fileName, n, currentPath)
	result <- log

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getStandardRings(subFolder, serverIP string, mode int) {
	// var wg sync.WaitGroup
	standardRingListName := "Ringlist.xml"
	distinctiveRingListName := "DistinctiveRingList.xml"
	var ringListName string
	results := make(chan string, 10) // we’ll write results into the buffered channel of strings

	perm := os.FileMode(0777)
	err := os.Mkdir(filepath.Join("./", subFolder), perm)
	check(err)

	switch mode {
	case 1:
		ringListName = standardRingListName
	case 2:
		ringListName = distinctiveRingListName
	default:
		fmt.Println("Invalid mode selected")
		os.Exit(1)
	}
	go getTftpFile(ringListName, subFolder, serverIP, results)
	path2 := filepath.Join("./", subFolder, ringListName)
	result := <-results
	rings := readRingListFile(path2)
	fmt.Println(result)
	for _, ring := range rings.Ring {
		// wg.Add(1)
		go getTftpFile(ring.FileName, subFolder, serverIP, results)
		// wg.Done()
	}
	for i := 0; i < len(rings.Ring); i++ {
		res := <-results
		fmt.Println(res)
	}
	// wg.Wait()
	close(results)
}

func createFile(path string) (*os.File, error) {
	return os.Create(path)
}
func writeToLog(file *os.File, msg string, offSet int64) error {
	bs, err := json.Marshal(&msg)
	check(err)
	_, err = file.WriteAt(bs, offSet)
	return err
}
