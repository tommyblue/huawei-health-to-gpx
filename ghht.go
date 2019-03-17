package ghht

import (
	"log"
	"strconv"
)

type GHHT struct {
	DbPath    string
	FileIndex int
}

func Init(dbPath string, fileIndex string) *GHHT {
	i, err := strconv.Atoi(fileIndex)

	if err != nil {
		log.Fatal(err)
	}
	mainConf := &GHHT{
		DbPath:    dbPath,
		FileIndex: i,
	}
	return mainConf
}
