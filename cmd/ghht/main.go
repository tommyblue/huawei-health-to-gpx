package main

import (
	"fmt"
	"os"

	ghht "github.com/tommyblue/go-huawei-health-tcx"
	"github.com/tommyblue/go-huawei-health-tcx/db"
	"github.com/tommyblue/go-huawei-health-tcx/tcx"
)

func main() {
	conf := ghht.Init(os.Args[1])

	database := db.GetDb(conf)
	defer database.Close()

	tracks := database.GetTracks()

	for _, t := range tracks {
		tcxTrack := tcx.FromDump(t)

		fmt.Println(tcxTrack)
	}
}
