package main

import (
	"log"
	"os"

	hitrack2gpx "github.com/tommyblue/huawei-health-to-gpx"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing DB Path")
	}
	i := "0"
	if len(os.Args) >= 3 {
		i = os.Args[2]
	}
	conf := hitrack2gpx.Init(os.Args[1], i)

	database := hitrack2gpx.GetDb(conf)
	defer database.Close()

	trackDump := database.GetTracks(conf.FileIndex)

	if i != "0" {
		track := hitrack2gpx.ParseTrackDump(trackDump)

		hitrack2gpx.GPXFromDump(track)
	}
}
