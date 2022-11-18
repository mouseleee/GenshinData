package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const driver = "sqlite3"

var (
	pool *sql.DB
	err  error
)

func InitConn(db string) {
	if _, err := os.Stat(db); err != nil {
		err := os.MkdirAll(filepath.Dir(db), os.ModeDir|0o700)
		if err != nil {
			fmt.Fprintf(os.Stderr, "创建数据库文件失败：%v\n", err)
		}
	}

	pool, err = sql.Open(driver, db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "连接数据库失败：%v\n", err)
	}

	if err := pool.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "访问数据库失败：%v\n", err)
	}
}

func CloseConn() {
	if err := pool.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "关闭数据库失败：%v\n", err)
	}
}

// func InitDb() {
// 	roles := `CREATE TABLE IF NOT EXISTS roles_index (
// 		id integer PRIMARY KEY AUTOINCREMENT,
// 		nickname varchar(10) NOT NULL,
// 		url varchar(100) NOT NULL
// 	)`
// }
