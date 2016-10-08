package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/jschweizer78/ciscoSsh/models"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	// "google.golang.org/api/drive/v2"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/sheets/v4"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("drive-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func checkErr(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s %v", msg, err))
	}
}

func main() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	checkErr("Unable to read client secret file:", err)

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/drive-go-quickstart.json
	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope, drive.DriveFileScope, drive.DriveScope)
	checkErr("Unable to parse client secret file to config:", err)
	client := getClient(ctx, config)

	// driveSrv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}
	// TODO Make TUI to get values
	var mode string
	fmt.Println("Would you like demo mode y/n")
	fmt.Scanln(&mode)

	sheetSrv, err := sheets.New(client)
	checkErr("Unable to retrieve Sheets Client", err)
	var syncStruct SyncStruct
	syncStruct.DoneChan = make(chan bool, 5)
	syncStruct.PrintChan = make(chan string, 5)
	syncStruct.LogChan = make(chan string, 10)
	as := getAppConfig(mode, &syncStruct)
	// readRange := fmt.Sprint(site, "!", dataRange)

	go func() {
		for {
			logMsg := <-as.channels.LogChan
			path := filepath.Join(as.Dir, "log.txt")
			file, err := os.OpenFile(path, os.O_APPEND, 0660)
			checkErr("Error opening log file", err)
			defer file.Close()

			writer := bufio.NewWriter(file)
			log := fmt.Sprintf("%s\n", logMsg)

			_, err = writer.WriteString(log)
			checkErr("Couldn't write to log file", err)
		}
	}()

	// TODO Just playing aroung
	// fileCall := driveSrv.Files.Get(spreadsheetID)
	// parents, err := fileCall.Fields("parents").Do()
	// checkErr("Problem with call", err)
	// fmt.Println(parents)

	sSheet := models.NewSheetTable(sheetSrv)
	as.st = &sSheet
	sSheet.SetDataLocation(as.SpreadsheetID, as.Site, as.DataRange, as.OffSet)
	fmt.Println(as.DataRange)
	// sSheet.LoadData()

	var wgWrite sync.WaitGroup
	go func() {
		// TODO get a list of device types and print to channel

		// as.channels.PrintChan <- dt
	}()
	go func() {
		for {
			deviceType, more := <-as.channels.PrintChan
			if more {
				wgWrite.Add(1)
				go writeCsvFileFromRespValues(deviceType, as, wgWrite)
			} else {
				as.channels.LogChan <- fmt.Sprintln("received all jobs")
				as.channels.DoneChan <- true
				return
			}
		}
	}()

	fmt.Println("Data loaded")
	// for i, aSw := range as.st.Switches {
	// 	fmt.Printf("Name: %s\tIP: %s\tRow: %d\n", aSw.Hostname, aSw.IPAddress, i+9)
	// }

	<-as.channels.DoneChan
	fmt.Println("Done requesting file to be written")
	wgWrite.Wait()
	fmt.Println("Done writting files")
	fmt.Println("Did it work?")

	// resp, err := sheetSrv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//

	// r, err := srv.Files.List().PageSize(10).
	// 	Fields("nextPageToken, files(id, name)").Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve files: %v", err)
	// }
	//
	// fmt.Println("Files:")
	// if len(r.Files) > 0 {
	// 	for _, i := range r.Files {
	// 		fmt.Printf("%s (%s)\n", i.Name, i.Id)
	// 	}
	// } else {
	// 	fmt.Println("No files found.")
	// }

}
