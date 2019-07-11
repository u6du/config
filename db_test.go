package config

import (
	"os"
	"testing"

	"github.com/u6du/ex"
)

func TestDb(t *testing.T) {
	name := "test/db"
	db := Db(
		name,

		`CREATE TABLE "dot" (
"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
"host"	TEXT NOT NULL UNIQUE,
"delay"	INTEGER NOT NULL DEFAULT 0);
CREATE INDEX "dot.delay" ON "dot" ("delay" ASC);`,

		"INSERT INTO dot(host) values (?)",

		"dns.rubyfish.cn",
		"dot-jp.blahdns.com",
	)

	li := db.Query("select id,host from dot")

	var id uint64
	var host string
	count := 0
	for li.Next() {
		ex.Warn(li.Scan(&id, &host))
		//t.Log(id, host)
		count++
	}
	if 0 == count {
		t.Error("row count = 0")
	}
	ex.Warn(os.Remove(File.Path(name + ".sqlite3")))
}
