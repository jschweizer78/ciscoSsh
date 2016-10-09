package models

import (
	"fmt"
	"sync"

	// storage "github.com/jschweizer78/ciscoSsh/storage/modelStorage"
	"google.golang.org/api/sheets/v4"
)

// SheetTable is a model of a Backbone information sheet site switch
type SheetTable struct {
	Headers  map[string]Header
	Switches []Switch
	// Storage   storage.STInterface
	sheetID   string
	site      string
	dCols     string
	headerRow int
	sServ     *sheets.Service
	Resp      [][]interface{}
	RespData  [][]string
}

// Header is a model of any header and where it is on the sheet
type Header struct {
	Name string
	Val  string
	Pos  int
}

// NewSheetTable Used to create a new sheet table with map/slice ready for use
func NewSheetTable(sServ *sheets.Service) SheetTable {
	return SheetTable{Headers: make(map[string]Header), Switches: make([]Switch, 50), sServ: sServ}
}

// SetDataLocation loads the switches into this SwitchSheet
func (st *SheetTable) SetDataLocation(sheetID, site, dCols string, pos int) {
	st.sheetID = sheetID
	st.site = site
	st.dCols = dCols
	st.headerRow = pos
}

// MakeRangeString is used to return a string of data range with header column set as start of the set.
func (st *SheetTable) MakeRangeString() string {
	return fmt.Sprint(st.dCols)
}

// LoadData gets the data from google Spreadsheet and looads it into the structs
func (st *SheetTable) LoadData() {
	var wg sync.WaitGroup
	resp, err := st.sServ.Spreadsheets.Values.Get(st.sheetID, st.dCols).Do()
	wg.Add(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wg.Done()
	if len(resp.Values) > 0 {
		st.Resp = resp.Values
		for i, row := range resp.Values {
			if i == 0 {
				st.Headers["Hostname"] = Header{Name: "Hostname", Val: fmt.Sprint(row[3]), Pos: 3}
				st.Headers["IDF"] = Header{Name: "IDF", Val: fmt.Sprint(row[0]), Pos: 0}
				st.Headers["IPAddress"] = Header{Name: "IPAddress", Val: fmt.Sprint(row[1]), Pos: 1}
				st.Headers["Model"] = Header{Name: "Model", Val: fmt.Sprint(row[8]), Pos: 8}
				st.Headers["Room"] = Header{Name: "Room", Val: fmt.Sprint(row[2]), Pos: 2}
			} else {
				for j, cell := range row {
					switch j {
					case 0:
						st.Switches = append(st.Switches, Switch{IDF: fmt.Sprint(cell)})
					case 1:
						st.Switches[i-1].IPAddress = fmt.Sprint(cell)
					case 2:
						st.Switches[i-1].Room = fmt.Sprint(cell)
					case 3:
						st.Switches[i-1].Hostname = fmt.Sprint(cell)
					case 8:
						st.Switches[i-1].Model = fmt.Sprint(cell)
					}
				}

			}

		}
	} else {
		fmt.Print("No data found.")
	}

}

// ConvertRespValues is used to convert []interface{} to []string
func (st *SheetTable) ConvertRespValues() {
	rv := st.Resp
	converted := make([][]string, len(rv))
	for i := range converted {
		converted[i] = make([]string, len(rv[0]))
	}
	for i, valSlice := range rv {
		for j, cell := range valSlice {
			converted[i][j] = fmt.Sprint(cell)
		}
	}
	st.RespData = converted
}

func convertRespValuesTest(rv [][]interface{}) [][]string {
	converted := make([][]string, len(rv))
	var wgBuild sync.WaitGroup
	var wgWrite sync.WaitGroup
	for pos := range converted {
		wgBuild.Add(1)
		go func(pos int) {
			converted[pos] = make([]string, len(rv[0]))
			wgBuild.Done()
		}(pos)
	}
	wgBuild.Wait()
	for i, valSlice := range rv {
		wgWrite.Add(1)
		go func(row int, vals []interface{}) {
			for j, cell := range vals {
				go func(col int, data interface{}) {
					converted[row][col] = fmt.Sprint(data)
					wgWrite.Done()
				}(j, cell)
			}
		}(i, valSlice)
	}
	wgWrite.Wait()
	return converted
}
