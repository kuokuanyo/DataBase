package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//檢查是否錯誤
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//初始化並連接資料庫function
//完整的資料格式連線: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func Connection(DriverName, user, password, host, database string) *sql.DB {

	//資料庫連結
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		user, password, host, database)

	//開啟資料庫連線(sql.Open只是初始化sql.DB物件)
	//func Open(driverName, dataSourceName string) (*DB, error)
	//第一個參數為驅動名稱，第二個參數為資料庫的連結
	db, err := sql.Open(DriverName, DataSourceName)
	//檢查錯誤
	checkError(err)
	//defer
	//defer db.Close()

	//立即檢查資料庫連線是否可用
	//func (db *DB) Ping() error
	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	return db
}

//新建資料庫
func CreateDb(db *sql.DB, DbName string) {

	CheckDb := fmt.Sprintf("DROP DATABASE IF exists %s", DbName)

	//檢查是否有存在的DATABASE
	//如果有，刪除
	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err := db.Exec(CheckDb)

	CerateName := fmt.Sprintf("CREATE DATABASE %s", DbName)

	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err = db.Exec(CerateName)
	checkError(err)

	//使用資料庫
	UseName := fmt.Sprintf("USE %s", DbName)

	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err = db.Exec(UseName)
	checkError(err)
	fmt.Println("Successfully created database")
}

//新建table
func CreateTable(db *sql.DB, TableName string, col_name []string, col_type []string) {

	//檢查是否有存在的table
	CheckTable := fmt.Sprintf("DROP TABLE IF exists %s", TableName)

	//如果有，刪除
	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err := db.Exec(CheckTable)
	checkError(err)

	//建立新table
	//檢查長度是否符合
	if len(col_name) != len(col_type) {
		log.Fatal("Parameters length not true")
	}

	//欄位長度
	n := len(col_name)

	CreateTable := fmt.Sprintf("CREATE TABLE %s (", TableName)

	//加入欄位及類型
	for i := 0; i < n; i++ {
		if i == n-1 {
			CreateTable += fmt.Sprintf("%s %s)", col_name[i], col_type[i])
		} else {
			CreateTable += fmt.Sprintf("%s %s, ", col_name[i], col_type[i])
		}
	}

	//func (db *DB) Exec(query string, args ...interface{}) (Result, error)
	_, err = db.Exec(CreateTable)
	//檢查錯誤
	checkError(err)

	fmt.Println("Successfully created table")
}

//插入值
func Insert_value(db *sql.DB, TableName string, col_name []string, args ...interface{}) {

	//加入欄位的長度
	n := len(col_name)

	//檢查欄位個數是否與加入值的個數符合
	if len(col_name) != len(args) {
		log.Fatal("length not true.")
	}

	Insert_Col := fmt.Sprintf("INSERT INTO %s(", TableName)

	//加入欄位
	for i := 0; i < n; i++ {
		if i == n-1 {
			Insert_Col += fmt.Sprintf("%s) values (", col_name[i])
		} else {
			Insert_Col += fmt.Sprintf("%s, ", col_name[i])
		}
	}

	//新增占位符
	for i := 0; i < n; i++ {
		if i == n-1 {
			Insert_Col += "?)"
		} else {
			Insert_Col += "?, "
		}
	}

	_, err := db.Exec(Insert_Col, args...)
	checkError(err)

	fmt.Println("Successfully insert value.")
}

//更新資料
func Update(db *sql.DB, TableName string, set_name string, where_name string, args ...interface{}) {

	//字串
	Update_str := fmt.Sprintf(`UPDATE %s set %s = `, TableName, set_name)

	//判別set_value數值類型
	switch args[0].(type) {
	case int:
		Update_str += fmt.Sprintf("%d where %s = ", args[0], where_name)
	case string:
		Update_str += fmt.Sprintf(`"%s" where %s = `, args[0], where_name)
	}

	//判別args[1]數值類型
	switch args[1].(type) {
	case int:
		Update_str += fmt.Sprintf("%d", args[1])
	case string:
		Update_str += fmt.Sprintf(`"%s"`, args[1])
	}

	_, err := db.Exec(Update_str)
	checkError(err)

	fmt.Println("Successfully Updated.")
}

//刪除資料
func Delete_data(db *sql.DB, TableName string, where_name string, args ...interface{}) {

	//字串
	Delete_str := fmt.Sprintf("DELETE from %s where %s = ?", TableName, where_name)

	//func (db *DB) Prepare(query string) (*Stmt, error)
	//Prepaer method 是為了之後的運行先準備好的語法
	//如需要運行此語法，執行stmt.Exec()
	//可以同時運行多個查詢
	stmt, err := db.Prepare(Delete_str)
	//執行
	stmt.Exec(args[0])
	//關閉stmt
	defer stmt.Close()
	//檢查錯誤
	checkError(err)

	fmt.Println("Successfully Deleted.")
}

//刪除資料庫
func Delete_Db(db *sql.DB, DbName string) {

	//刪除字串
	Delete := fmt.Sprintf("DROP DATABASE %s", DbName)

	//刪除
	db.Exec(Delete)

	fmt.Println("Successfully Deleted Database.")
}

//刪除資料表
func Delete_Tb(db *sql.DB, TableName string) {

	//刪除字串
	Delete := fmt.Sprintf("DROP TABLE %s", TableName)

	//刪除
	db.Exec(Delete)

	fmt.Println("Successfully Deleted Table.")
}

//讀取資料
func Read_Table(db *sql.DB, TableName string, args ...interface{}) {

	//讀取字串
	//"SELECT col from tablename;"
	Read_str := "SELECT "

	//args長度
	n := len(args)

	//添加字串
	for i := 0; i < n; i++ {
		if i == n-1 {
			Read_str += fmt.Sprintf("%s from %s;", args[i], TableName)
		} else {
			Read_str += fmt.Sprintf("%s, ", args[i])
		}
	}

	//讀取
	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	rows, err := db.Query(Read_str)
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
		if n == 1 {
			err := rows.Scan(&args[0])
			checkError(err)
		} else if n == 2 {
			err := rows.Scan(&args[0], &args[1])
			checkError(err)
		} else if n == 2 {
			err := rows.Scan(&args[0], &args[1], &args[2])
			checkError(err)
		}
		fmt.Printf("Data row = %s\n", args...)
	}

	//在迴圈中是否有錯誤
	err = rows.Err()
	checkError(err)
	fmt.Println("Read Done.")
}

func main() {

	//設定需要添加欄位
	col_name := []string{"id", "name", "quantity"}
	col_type := []string{"int", "VARCHAR(50)", "int"}

	//test_col := []string{"id", "name"}
	//連接資料庫
	db := Connection("mysql", "root", "asdf4440", "127.0.0.1", "test")
	//defer(關閉資料庫)
	defer db.Close()

	//新建資料庫
	CreateDb(db, "t")

	//新建資料表
	CreateTable(db, "newtest", col_name, col_type)
	/*
		//加入值
		Insert_value(db, "newtest", col_name, 1, "John", 45)
		Insert_value(db, "newtest", test_col, 3, "David")

		//更新
		Update(db, "newtest", "quantity", "name", 100, "David")
		Update(db, "newtest", "name", "quantity", "David", 100)

		//刪除值
		Delete_data(db, "newtest", "name", "David")

		//讀取
		Read_Table(db, "newtest", "name")
	*/
}
