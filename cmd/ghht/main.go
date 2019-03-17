package main

import (
	"fmt"
	"log"
	"os"

	ghht "github.com/tommyblue/go-huawei-health-tcx"
	"github.com/tommyblue/go-huawei-health-tcx/db"
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

	database := db.GetDb(conf)
	defer database.Close()

	track := database.GetTracks(conf.FileIndex)
	fmt.Println(track)
	// tracks := database.GetTracks(conf.FileIndex)

	// for _, t := range tracks {
	// 	tcxTrack := tcx.FromDump(t)

	// 	fmt.Println(tcxTrack)
	// }
}
