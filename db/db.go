package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

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

func (db *DB) GetTracks(fileIndex int) string {
	// var acc []string
	files := db.getFiles()
	if fileIndex == 0 {
		fmt.Println("\nSelect an ID from the list above and pass it as second argument\n")
		return ""
	}
	for _, id := range files {
		if id == fileIndex {
			return db.getTrack(id)
		}
		// acc = append(acc, db.getTrack(id))
	}
	// return acc
	log.Fatal("Cannot find the selected ID")
	return ""
}

func (db *DB) getFiles() []int {
	query := `SELECT file_index, file_path FROM apk_file_info WHERE file_path LIKE '%HiTrack%';`

	var acc []int

	callback := func(scanFn scannerFn) {
		var fileIndex int
		var filePath string
		err := scanFn(&fileIndex, &filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("File %s (ID %d)\n", filePath, fileIndex)
		acc = append(acc, fileIndex)
	}

	db.makeQuery(query, callback)

	return acc
}

func (db *DB) getTrack(id int) string {
	query := fmt.Sprintf(`SELECT file_data FROM apk_file_data WHERE file_index=%d ORDER BY data_index;`, id)
	// lines to be joined. If doesn't end with ; is interrupted?
	var b bytes.Buffer
	callback := func(scanFn scannerFn) {
		var fileData string
		err := scanFn(&fileData)
		if err != nil {
			log.Fatal(err)
		}
		b.WriteString(fileData)
	}

	db.makeQuery(query, callback)
	return b.String()
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
