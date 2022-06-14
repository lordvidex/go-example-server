package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	// load JSON data from data/mock.json
	err := loadData()
	if err != nil {
		panic("An error occured loading json data: " + err.Error())
	}
}

func loadData() (err error) {
	var byteValue []byte
	file, err := os.Open(filepath.Join("data", "mock.json")) // open file
	if err != nil {
		goto END
	}

	byteValue, err = ioutil.ReadAll(file) // read bytes from file
	if err != nil {
		goto END
	}

	err = json.Unmarshal(byteValue, &products) // read to json
	if err != nil {
		goto END
	}
END:
	return
}