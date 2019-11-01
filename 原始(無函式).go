//連接資料庫套件
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	database = "test"
	user     = "root"
	password = "asdf4440"
)

//檢查是否有錯誤
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//完整的資料格式連線: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func main() {

	//初始化資料庫連線
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		user, password, host, database)

	//開啟資料庫連線(sql.Open只是初始化sql.DB物件)
	//func Open(driverName, dataSourceName string) (*DB, error)
	//第一個參數為驅動名稱，第二個參數為資料庫的連結
	db, err := sql.Open("mysql", connectionString)
	//檢查錯誤
	checkError(err)
	//defer
	defer db.Close()

	//立即檢查資料庫連線是否可用
	//func (db *DB) Ping() error
	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	//建立想要查詢資料欄位變數
	var (
		name         string
		id, quantity int
	)

	/*建立新資料庫
	_, err = db.Exec("create database testDB")
	checkError(err)
	//使用資料庫
	_,err = db.Exec(“USE testDB”)
	checkError(err)
	*/

	//檢查是否有存在的table
	//如果有，刪除
	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err = db.Exec("DROP TABLE IF exists newtable")
	checkError(err)
	fmt.Println("Finished dropping table(if exists.)")

	//建立新table
	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err = db.Exec("CREATE TABLE newtable(id integer, name varchar(50),quantity integer)")
	//檢查錯誤
	checkError(err)
	fmt.Println("Successfully created table")

	//插入值
	//func (db *DB) Prepare(query string) (*Stmt, error)
	//Prepaer method 是為了之後的運行先準備好的語法
	//如需要運行此語法，執行stmt.Exec()
	//可以同時運行多個查詢
	stmt, err := db.Prepare("INSERT INTO newtable (id, name, quantity) values (?, ?, ?)")
	//使用後需要關閉
	defer stmt.Close()
	checkError(err)
	//運行Prepare method
	//加入值
	_, err = stmt.Exec(1, "b", 100)
	checkError(err)
	_, err = stmt.Exec(2, "b", 100)
	checkError(err)

	//更新資料
	//func (db *DB) Prepare(query string) (*Stmt, error)
	_, err = db.Exec(`UPDATE newtable set quantity = 100 where name = "b"`)
	checkError(err)

	//刪除值
	//func (db *DB) Prepare(query string) (*Stmt, error)
	_, err = db.Exec("DELETE from newtable where id = 1")
	checkError(err)

	//讀取資料
	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	rows, err := db.Query("SELECT id, name, quantity from newtable;")
	//檢查錯誤
	checkError(err)
	//defer 關閉查詢
	defer rows.Close()
	fmt.Println("Reading Data:")

	//Next method 迭代查詢資料，回傳bool
	//func (rs *Rows) Next() bool
	for rows.Next() {

		//Scan method方法用來讀取每一列的值
		//func (rs *Rows) Scan(dest ...interface{}) error
		err := rows.Scan(&id, &name, &quantity)
		checkError(err)
		fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
	}
	//在迴圈中是否有錯誤
	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")
}
