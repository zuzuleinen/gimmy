package infrastructure

type DB struct {
	Counter int
}

var d *DB

func GetDB() *DB {
	if d == nil {
		d = &DB{}
	}
	return d
}
