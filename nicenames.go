package main

import (
	"os"
	"log"
	"json"
	"strings"
)

type niceName struct {
	Name string
	ID byte
}

var niceNames []niceName
var loadedNiceNames = false

func LoadNiceNames() {
	f, err := os.Open("nicenames.json", os.O_RDONLY, 0644)
	if err != nil {
		// Fatal until i add proper error handling
		log.Fatal(err)
		return
	}
	decoder := json.NewDecoder(f)
	niceNames = make([]niceName, 256)
	err = decoder.Decode(&niceNames)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (b *Block) NiceName() string {
	if !loadedNiceNames {
		LoadNiceNames()
	}
	for _, niceName := range niceNames {
		if niceName.ID == b.Type {
			return strings.Title(niceName.Name)
		}
	}
	return "Unknown"
}