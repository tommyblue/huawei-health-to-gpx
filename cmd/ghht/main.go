package main

import (
	"os"

	ghht "github.com/tommyblue/go-huawei-health-tcx"
	"github.com/tommyblue/go-huawei-health-tcx/db"
)

func main() {
	conf := ghht.Init(os.Args[1])
	database := db.GetDb(conf)
	defer database.Close()

	db.GetFiles(database)
}
