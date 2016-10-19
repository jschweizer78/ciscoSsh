package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	// "sync"

	"github.com/jschweizer78/ciscoSsh/models"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	// "google.golang.org/api/drive/v2"

	"google.golang.org/api/drive/v2"
	"google.golang.org/api/sheets/v4"
)

var dataReadyChanName = `DATA IS READY`

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
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve drive Client %v", err)
	// }
	// TODO Make TUI to get values
	var mode string
	fmt.Println("Would you like demo mode y/n")
	fmt.Scanln(&mode)

	sheetSrv, err := sheets.New(client)
	checkErr("Unable to retrieve Sheets Client", err)
	syncer := newSyscStruct()
	as := getAppConfig(mode, syncer)
	// as.wgs[`wfDATA`]

	// TODO Just playing aroung
	// fileCall := driveSrv.Files.Get(spreadsheetID)
	// parents, err := fileCall.Fields("parents").Do()
	// checkErr("Problem with call", err)
	// fmt.Println(parents)
	// myDbName := "MyDatabase"
	// myDB := tiedot.NewStorageDB(filepath.Join(as.Dir, myDbName))
	// userStor := tiedot.NewUserStorage(myDB)
	sSheet := models.NewSheetTable(sheetSrv)
	sSheet.SetDataLocation(as.SpreadsheetID, as.Site, as.DataRange, as.OffSet)
	fmt.Println(as.DataRange)
	// sSheet.LoadData()
	sSheet.ConvertRespValues()
	as.syncer.FilesChan[dataReadyChanName] = make(chan string, 0)
	as.St = &sSheet

	fmt.Println("Data loaded")
	ws := newWebServer()
	// userRes := resource.UserResource{Stor: tiedot.NewUserStorage(myDB)}
	// ws.api[userRes.MyName()] = userRes
	ws.run()

	// for i, aSw := range as.st.Switches {
	// 	fmt.Printf("Name: %s\tIP: %s\tRow: %d\n", aSw.Hostname, aSw.IPAddress, i+9)
	// }

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
	fmt.Println("Did it work?")

}
