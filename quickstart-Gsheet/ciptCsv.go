package main

import (
	"fmt"
	"sync"
)

var headerPOS = 13

func ciptData() {
	voipImportSpreadSheetID := "1p-X_yc9Rsu97F3QLnUcA-BtG2061ELGAbHUGr4BLcOw"
	sourceSheet := "Sites"
	cutSheet := "Current Cut Sheet - ReadOnly"
	lastCol := "AK"
	fmt.Printf("%s \n%s!:%s From SpreadsheetID:%s", sourceSheet, cutSheet, lastCol, voipImportSpreadSheetID)
}

func writeCsvFileFromRespValues(devType string, ac *AppConfig, wg sync.WaitGroup) {
	// TODO Need to find device type header and remove column from datSet
	data := convertRespValues(ac.st.Resp)
	data = removeDevHeaderFromData(data)
	switch devType {
	case "CTI Port", "Jabber", "Cisco 8831", "Analog":
		go makeSingleLinePhone(devType, data, wg)
	case "Cisco 8851", "Cisco 8841", "Cisco 7841":
		go makeMultiLinePhone(devType, data, wg)
	default:
		ac.channels.LogChan <- fmt.Sprintf("Unknown device type:%s", devType)
	}

}

func convertRespValues(rv [][]interface{}) [][]string {
	converted := make([][]string, len(rv))
	for i := range converted {
		converted[i] = make([]string, len(rv[0]))
	}
	for i, valSlice := range rv {
		for j, cell := range valSlice {
			converted[i][j] = fmt.Sprint(cell)
		}
	}
	return converted
}

func filterDeviceFromData(dt string, data [][]string) [][]string {
	filtered := make([][]string, len(data))
	for i := range filtered {
		filtered[i] = make([]string, len(data[i]))
	}
	for _, item := range data {
		// TODO make the device header position a variable
		if item[headerPOS] == dt {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
func removeDevHeaderFromData(data [][]string) [][]string {
	for i, item := range data {
		data[i] = append(item[0:headerPOS], item[headerPOS+1:]...)
	}
	return data
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
