package main

import (
	"log"
	"os"

	ghht "github.com/tommyblue/go-huawei-health-tcx"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing DB Path")
	}
	i := "0"
	if len(os.Args) >= 3 {
		i = os.Args[2]
	}
	conf := ghht.Init(os.Args[1], i)

	database := ghht.GetDb(conf)
	defer database.Close()

	trackDump := database.GetTracks(conf.FileIndex)
	// fmt.Println(trackDump)
	track := ghht.ParseTrackDump(trackDump)

	ghht.GPXFromDump(track)
	// fmt.Printf("\n%v\n", gpxDoc)
	// tracks := database.GetTracks(conf.FileIndex)

	// for _, t := range tracks {
	// 	tcxTrack := tcx.FromDump(t)

	// 	fmt.Println(tcxTrack)
	// }
}
