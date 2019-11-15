package main

import (
	"conn"
)

//設定資料庫資訊
var user = conn.MySqlUser{
	Host:      //主機
	MaxIdle:   //閒置的連接數
	MaxOpen:   //最大連接數
	User:      //用戶名
	Password:  //密碼
	Database:  //資料庫名稱
	Port:      //端口
}

//建立查詢欄位
var (
	名稱	int
	名稱	string
	名稱	bool
)

//上面查詢欄位名稱等於此[]string{}的變數名稱
//須為字串
var s = []string{上列設定名稱}

func main() {

	//建立初始化連線
	connect_db := user.Init()

	//回傳指標
	db = conn.NewDB(connect_db)

	//最後必須關閉
	defer db.Close()

	//建立資料庫
	db.CreateDb(資料庫名稱)

	//使用資料庫
	db.Use_Db(資料庫名稱)

	//建立資料表
	db.CreateTable(資料表名稱, 欄位名稱, 欄位類型...)

	//插入數值
	db.Insert(資料表明撐, 插入欄位名稱, 插入數值...)

	//更改數值
	db.Update_db(資料庫名稱, 設定欄位名稱, 設定新數值, 更改的欄位, 更改欄位的數值)

	//刪除資料庫
	db.Delete_Db(資料庫名稱)

	//刪除資料表
	db.Delete_Tb(資料庫名稱)

	//讀取資料
	//第三個與後面參數長度必須相同
	db.Read(資料庫名稱, s, 設定的變數(var))

}
