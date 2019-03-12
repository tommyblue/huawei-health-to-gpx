package main

func main() {
	db := getDb()
	defer db.Close()

	getFiles(db)
}
