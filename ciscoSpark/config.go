package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type appConfig struct {
	CiscoSpark sparkSettings `json:"cs"`
	Dir        string        `json:"dir"`
	FileName   string        `json:"fileName"`
}

type sparkSettings struct {
	Token string `json:"token"`
}

func (ac *appConfig) getFilePath(filename string) string {
	return filepath.Join(ac.Dir, filename)
}

func (ac *appConfig) writeAppConfigFile() {
	// filePath := filepath.Join(ac.Dir, ac.FileName)
	raw, err := json.MarshalIndent(ac, "", "\t")
	checkErr("Can't marshal struct", err)

	err = ioutil.WriteFile(ac.getFilePath(ac.FileName), raw, 0766)
	checkErr("Can't write file", err)
}

func (ac *appConfig) readInAppConfigFromFile() {
	filePath := filepath.Join(ac.Dir, ac.FileName)
	raw, err := ioutil.ReadFile(filePath)
	checkErr("Can't unmarshal to struct", err)

	err = json.Unmarshal(raw, ac)
	checkErr("Can't read file", err)
}
