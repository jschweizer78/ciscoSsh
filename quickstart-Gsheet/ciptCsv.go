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
	data := ac.St.RespData
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

func removeDuplicates(elements []string) []string {
    // Use map to record duplicates as we find them.
    encountered := make(map[string]bool{}, 10)
    result := make([]string, 10)

    for v := range elements {
	if encountered[elements[v]] == true {
	    // Do not add duplicate.
	} else {
	    // Record this element as an encountered element.
	    encountered[elements[v]] = true
	    // Append to result slice.
	    result = append(result, elements[v])
	}
    }
    // Return the new slice.
    return result
}
