package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	ghht "github.com/tommyblue/go-huawei-health-tcx"

	_ "github.com/mattn/go-sqlite3"
)

func GetDb(conf *ghht.GHHT) *sql.DB {
	db, err := sql.Open("sqlite3", conf.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetFiles(db *sql.DB) {
	query := `SELECT file_index, file_path FROM apk_file_info WHERE file_path LIKE '%HiTrack%';`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var fileIndex int
		var filePath string
		err = rows.Scan(&fileIndex, &filePath)
		if err != nil {
			log.Fatal(err)
		}
		paths := strings.Split(filePath, "/")
		fmt.Println(fileIndex, paths[len(paths)-1])
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func GetTrack() {
	// q2 := `SELECT file_data FROM apk_file_data WHERE file_index=11 ORDER BY data_index;`
	// lines to be joined. If doesn't end with ; is interrupted?
}
