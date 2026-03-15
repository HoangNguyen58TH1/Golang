package my_sql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() {
	db := ConnectDB()
	InsertUser(db, "Hoang", 50, "Software Engineer", "Male")
	GetUsers(db)
	DeleteUserByName(db, "hoang")
	GetUsers(db)
}

func ConnectDB() *sql.DB {
	// dsn = data source name = contain all infos to connect DB
	// dsn := "username:pw@tcp(DB_host:MySQL_port)/DB"
	dsn := "root:123456@tcp(localhost:3306)/golang"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// - sql.Open() không mở connection ngay.
	// - Connection thực sự chỉ được tạo khi: db.Query() / db.Exec() / db.Ping()
	// CHECK CONNECTION
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// SETTING CONNECTTION POOL
	db.SetMaxOpenConns(10)           // max connection to DB
	db.SetMaxIdleConns(5)            // Connection idle giữ trong pool (not use but keep to reuse)
	db.SetConnMaxLifetime(time.Hour) // Connection sống tối đa bao lâu, before recreate.

	fmt.Println("=== Connected to MySQL ===")
	return db
}

func InsertUser(db *sql.DB, name string, age int, job string, sex string) {
	insert_results, err := db.Exec(
		"INSERT INTO users (name, age, job, sex) VALUES (?, ?, ?, ?)",
		name, age, job, sex,
	)

	if err != nil {
		fmt.Println("insert error:", err)
		return
	}

	id, _ := insert_results.LastInsertId()
	fmt.Println("=== Inserted user id:", id)

	// DUPLICATE ID = 1 --> FAILED
	// _, err = db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", 1, "Hoang")
	// if err != nil {
	// 	panic(err)
	// }
}

func DeleteUserByName(db *sql.DB, name string) {
	delete_result, err := db.Exec("DELETE FROM users WHERE name = ?", name)

	if err != nil {
		panic(err)
	}

	rowsAffected, _ := delete_result.RowsAffected()
	fmt.Println("=== Rows deleted:", &rowsAffected)
}

func GetUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age, job, sex FROM users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var job string
		var sex string
		rows.Scan(&id, &name, &age, &job, &sex)
		fmt.Printf("--- Users info:  %d, %s, %d, %s, %s\n", id, name, age, job, sex)
	}
}
