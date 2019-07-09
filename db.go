package config

import (
	"database/sql"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	"github.com/u6du/ex"
)

const DriverName = "sqlite3"

func Db(name, create string, args ...interface{}) *sql.DB {
	name = name + "." + DriverName
	dbpath, isNew := File.PathIsNew(name)

	db, err := sql.Open(DriverName, dbpath)
	ex.Panic(err)

	if isNew {
		_, err := db.Exec(create)
		ex.Panic(err)

		argsLen := len(args)
		if argsLen > 0 {
			insert := args[0].(string)

			if argsLen > 1 {
				s, err := db.Prepare(insert)
				ex.Panic(err)

				for _, i := range args[1:] {
					t := reflect.TypeOf(i)
					switch t.Kind() {
					case reflect.Interface:
						li, _ := i.([]interface{})
						_, err = s.Exec(li...)
					default:
						_, err = s.Exec(i)
					}
					ex.Panic(err)

				}
			} else {
				_, err := db.Exec(insert)
				ex.Panic(err)
			}
		}
	}

	return db
}
