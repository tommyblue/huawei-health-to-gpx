package hitrack2gpx

import (
	"log"
	"strconv"
)

type HT2G struct {
	DbPath    string
	FileIndex int
}

func Init(dbPath string, fileIndex string) *HT2G {
	i, err := strconv.Atoi(fileIndex)

	if err != nil {
		log.Fatal(err)
	}
	mainConf := &HT2G{
		DbPath:    dbPath,
		FileIndex: i,
	}
	return mainConf
}
