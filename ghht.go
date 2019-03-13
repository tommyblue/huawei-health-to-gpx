package ghht

type GHHT struct {
	DbPath string
}

func Init(db_path string) *GHHT {
	mainConf := &GHHT{
		DbPath: db_path,
	}
	return mainConf
}
