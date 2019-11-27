package main

import (
	"conn"
	"os"

	"src/github.com/subosito/gotenv"
)

var db *conn.DB
var data conn.ColName
var datas []conn.ColName

func init() {
	gotenv.Load()
	//設定資料庫資訊
	user := conn.MySqlUser{
		Host:     os.Getenv("db_host"), //主機
		MaxIdle:  10,                   //閒置的連接數
		MaxOpen:  10,                   //最大連接數
		User:     os.Getenv("db_user"), //用戶名
		Password: os.Getenv("db_pass"), //密碼
		Database: os.Getenv("db_name"), //資料庫名稱
		Port:     os.Getenv("db_port"), //端口
	}

	db = user.Init()
}

func main() {
	//最後必須關閉
	defer db.Close()

	//建立資料庫
	db.CreateDb(資料庫名稱)

	//使用資料庫
	db.Use_Db(資料庫名稱)

	//建立資料表
	db.CreateTable(資料表名稱, 欄位名稱n, 欄位類型n...)

	//插入數值
	db.Insert(資料表名稱, 插入欄位名稱n, 插入數值n...)

	//更改數值
	db.Update_db(資料庫名稱, 設定欄位名稱, 設定新數值, 更改的欄位, 更改欄位的數值)

	//刪除資料庫
	db.Delete_Db(資料庫名稱)

	//刪除資料表
	db.Delete_Tb(資料庫名稱)

	//讀取資料
	db.ReadAll(資料庫名稱, datas, data)

	//讀取條件
	db.ReadOne(資料庫名稱, data, 查詢欄位名稱, 查詢欄位的值)
}
