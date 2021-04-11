package infrastructure

import (
	"github.com/google/uuid"
)

type DBConn struct {
	Tables map[string]Rows
}

type Rows map[string]Model

var d *DBConn

type Model interface {
	Key() string
	Data() []byte
}

func DB() *DBConn {
	if d == nil {
		d = &DBConn{}
		d.Tables = make(map[string]Rows)
	}
	return d
}

func (d *DBConn) Save(tableName string, c Model) {
	_, ok := d.Tables[tableName]
	if !ok {
		d.Tables[tableName] = make(map[string]Model)
	}
	d.Tables[tableName][c.Key()] = c
}

func GenerateID() string {
	return uuid.New().String()
}
