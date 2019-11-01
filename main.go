package main

import (
	"conn"
)

func main() {
	var user = conn.MySqlUser{
		Host:     "127.0.0.1", //主機
		MaxIdle:  10,          //閒置的連接數
		MaxOpen:  10,          //最大連接數
		User:     "root",      //用戶名
		Password: "asdf4440",  //密碼
		Database: "test",      //資料庫名稱
		Port:     3306,        //端口

	}

	//建立初始化連線
	db := user.Init()
	//最後必須關閉
	defer db.Close()

	//建立資料庫
	conn.CreateDb(db, "t")

	//使用資料庫
	conn.Use_Db(db, "t")

	//建立資料表
	conn.CreateTable(db, "aaa",
		"id", "integer",
		"name", "VARCHAR(50)",
		"quantity", "integer")

	//插入數值
	conn.Insert(db, "aaa", "id", 1, "name", "Kuo", "quantity", 100)
}
