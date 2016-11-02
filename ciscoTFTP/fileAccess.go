package main

import (
	"bytes"
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

	currentPath := filepath.Join("./", subFolder, fileName)
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
		panic(err.Error())
	}
}

func getStandardRings(subFolder, serverIP string, mode int) {
	// var wg sync.WaitGroup
	standardRingListName := "Ringlist.xml"
	distinctiveRingListName := "DistinctiveRingList.xml"
	signedStanRLName := "Ringlist.xml.sgn"
	signedDistRLName := "DistinctiveRingList.xml.sgn"
	var ringListName string
	results := make(chan string, 10) // we’ll write results into the buffered channel of strings

	perm := os.FileMode(0777)
	sigPath := filepath.Join("./", subFolder, "signed")
	unsigPath := filepath.Join("./", subFolder, "unsigned")
	err := os.MkdirAll(sigPath, perm)
	check(err)
	err = os.MkdirAll(unsigPath, perm)
	check(err)

	switch mode {
	case 1:
		ringListName = standardRingListName
	case 2:
		ringListName = distinctiveRingListName
	case 3:
		ringListName = signedStanRLName
	case 4:
		ringListName = signedDistRLName
	default:
		fmt.Println("Invalid mode selected")
		os.Exit(1)
	}

	go getTftpFile(ringListName, subFolder, serverIP, results)
	var buffer bytes.Buffer

	result := <-results
	buffer.WriteString(fmt.Sprintf("%s\n", result))

	rings := readRingListFile(filepath.Join("./", subFolder, ringListName))
	fmt.Println(result)
	for _, ring := range rings.Ring {
		// wg.Add(1)
		go getTftpFile(ring.FileName, unsigPath, serverIP, results)
		// wg.Done()
	}

	// if mode == 3 || mode == 4 {
	//
	// }

	for i := 0; i < len(rings.Ring); i++ {
		res := <-results
		fmt.Println(res)
		// logMsg := fmt.Sprintf("%s\n", res)
		buffer.WriteString(fmt.Sprintf("%s\n", res))
	}
	logFile, err := os.Create("./log.txt")
	check(err)
	logFile.WriteString(buffer.String())
	// wg.Wait()
	close(results)
}

func createFile(path string) (*os.File, error) {
	return os.Create(path)
}

// TODO make it so we can write to a log file whenever we need. Not just expect one set.
func writeToLog(file *os.File, msg string, offSet int64) error {
	bs, err := json.Marshal(&msg)
	check(err)
	_, err = file.WriteAt(bs, offSet)
	return err
}
