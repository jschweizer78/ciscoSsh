package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/jschweizer78/ciscoSsh/models"
)

// AppConfig is used to model the base configuration for the SSH application
type AppConfig struct {
	SpreadsheetID string `json:"spreadSheetId"`
	Site          string `json:"site"`
	FirstCol      string `json:"firstColumn"`
	LastCol       string `json:"lastColumn"`
	DataRange     string `json:"dataRange,omitempty"`
	OffSet        int    `json:"headerRow,omitempty"`
	Dir           string `json:"-"`
	FileName      string `json:"-"`
	PubKey        string `json:"pubKey,omitempty"`
	PrivateKey    string `json:"privateKey,omitempty"`
	channels      *SyncStruct
	St            *models.SheetTable
}

func getAppConfig(mode string, sync *SyncStruct) *AppConfig {

	appConfig := AppConfig{channels: sync}
	// TODO create function to check for pub/private key files and to create if needed

	bol, err := convertYN(mode)
	checkErr("Bad answer", err)
	switch bol {
	case true:
		runDemoMode(&appConfig)
	case false:
		runTerminalMode(&appConfig)
	}
	return &appConfig
}

func (ac *AppConfig) setAppConfigDefaults() {
	ac.Dir = "./config"
	ac.FileName = "config.json"

	ac.SpreadsheetID = "1k7vtPZfvLGWyaR61AyJhMa56H4DuDonMEf1K4lapQSk"
	ac.Site = "Site 1"
	ac.FirstCol = "A"
	ac.LastCol = "O"
	ac.OffSet = 8
	ac.DataRange = fmt.Sprintf("%s!%s%d:%s", ac.Site, ac.FirstCol, ac.OffSet, ac.LastCol)
	ac.PubKey = "pubkey.pem"
	ac.PrivateKey = "id_rsa"
	err := makeDir(ac.Dir)
	checkErr("Can't make dir", err)

}

func (ac *AppConfig) getFilePath(filename string) string {
	return filepath.Join(ac.Dir, filename)
}

func (ac *AppConfig) writeAppConfigFile() {
	// filePath := filepath.Join(ac.Dir, ac.FileName)
	raw, err := json.MarshalIndent(ac, "", "  ")
	checkErr("Can't marshal struct", err)

	err = ioutil.WriteFile(ac.getFilePath(ac.FileName), raw, 0700)
	checkErr("Can't write file", err)
}

func (ac *AppConfig) readAppConfigFile() {
	filePath := filepath.Join(ac.Dir, ac.FileName)
	raw, err := ioutil.ReadFile(filePath)
	checkErr("Can't unmarshal to struct", err)

	err = json.Unmarshal(raw, ac)
	checkErr("Can't read file", err)
}

func (ac *AppConfig) loadAppConfigFile() {
	if ac.Dir == "" {
		var dir string
		var mode string
		fmt.Println("Would you like to use the default dir (y/n)? (config)")
		fmt.Scanln(&mode)
		bol, err := convertYN(mode)
		checkErr("Bad answer", err)
		switch bol {
		case true:
			dir = "config"
		case false:
			fmt.Println("What would you like the directory name to be?")
			fmt.Scanln(&dir)
		}
		ac.Dir = filepath.Join("./", dir)
	}
	if ac.FileName == "" {
		var filename string
		var mode string
		fmt.Println("Would you like to use the default file name? (y/n)? (config.json)")
		fmt.Scanln(&mode)
		bol, err := convertYN(mode)
		checkErr("Bad answer", err)
		switch bol {
		case true:
			filename = "config.json"
		case false:
			fmt.Println("What would you like the file name to be?")
			fmt.Scanln(&filename)
		}
		ac.FileName = filename
	}
	fmt.Printf("The names are dir=%s file=%s", ac.Dir, ac.FileName)
	ac.readAppConfigFile()
}

func runDemoMode(ac *AppConfig) {
	ac.setAppConfigDefaults()
	var answer string
	pubKeyFilePath := ac.getFilePath(ac.PubKey)
	privKeyFilePath := ac.getFilePath(ac.PrivateKey)
	err := checkForFile(pubKeyFilePath)
	if err != nil {
		fmt.Printf("No public key file found at: %s\n", pubKeyFilePath)
		err = MakeSSHKeyPair(pubKeyFilePath, privKeyFilePath)
		checkErr("Couldn't make key pair", err)
	}
	err = checkForFile(privKeyFilePath)
	if err != nil {
		fmt.Printf("No private key file found at: %s\n", privKeyFilePath)
	}
	// fmt.Println("Please enter the username")
	// user := scanResponse()
	// if user == "" {
	// 	user = "jasonschweizer"
	// }
	// fmt.Println("Please enter your password.")
	// pw := scanResponse()
	// if pw == "" {
	// 	pw = "DONE"
	// }
	// hostAndPort := fmt.Sprintf("%s:%s", ac.st.Switches[0].IPAddress, "22")
	// client, session, err := connectToHost(user, hostAndPort, pw)
	// checkErr("Couldn't connect", err)
	// defer client.Close()
	// defer session.Close()
	fmt.Println("Do you want to save defaults? (y,n,)")
	fmt.Scanln(&answer)
	bol, err := convertYN(answer)
	checkErr("Bad answer", err)
	switch bol {
	case false:
		return
	case true:
		ac.writeAppConfigFile()
	}

}

func runTerminalMode(ac *AppConfig) {
	ac.loadAppConfigFile()
	if ac.DataRange == "" {
		ac.DataRange = fmt.Sprintf("%s!%s%d:%s", ac.Site, ac.FirstCol, ac.OffSet, ac.LastCol)
	}

	fmt.Println(ac.Site)
}

func convertYN(tf string) (bool, error) {
	var bol bool
	switch tf {
	case "y", "Y", "Yes":
		bol = true
	case "n", "N", "No", "":
		bol = false
	default:
		err := fmt.Errorf("Unkown answer %s", tf)
		return false, err
	}
	return bol, nil
}

func scanResponse() string {
	var word string
	_, err := fmt.Scanln(&word)
	checkErr("can't scan", err)
	return word
}

func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Not a number")
	}
	return i
}
