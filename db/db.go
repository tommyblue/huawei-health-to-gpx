package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	ghht "github.com/tommyblue/go-huawei-health-tcx"

	_ "github.com/mattn/go-sqlite3"
)

type scannerFn func(dest ...interface{}) error
type callbackFn func(scannerFn)

type DB struct {
	db *sql.DB
}

func GetDb(conf *ghht.GHHT) *DB {
	db, err := sql.Open("sqlite3", conf.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		db: db,
	}
}

func (db *DB) Close() {
	db.db.Close()
}

func (db *DB) GetFiles() {
	query := `SELECT file_index, file_path FROM apk_file_info WHERE file_path LIKE '%HiTrack%';`

	callback := func(scanFn scannerFn) {
		var fileIndex int
		var filePath string
		err := scanFn(&fileIndex, &filePath)
		if err != nil {
			log.Fatal(err)
		}
		paths := strings.Split(filePath, "/")
		fmt.Println(fileIndex, paths[len(paths)-1])
	}

	db.makeQuery(query, callback)

}

func (db *DB) GetTrack() {
	// q2 := `SELECT file_data FROM apk_file_data WHERE file_index=11 ORDER BY data_index;`
	// lines to be joined. If doesn't end with ; is interrupted?
}

func (db *DB) makeQuery(query string, callback callbackFn) {
	rows, err := db.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		callback(rows.Scan)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
