//讀取數據
package conn

import (
	"fmt"
	"log"
)

//查詢欄位名稱
type ColName struct {
	Name  string
	Area  string
	Total int
}

//function read all data
//SELECT col_name FROM tablename;
//args為scan的欄位名稱
func (db DB) ReadAll(TableName string, datas []ColName, data ColName) ([]ColName, error) {

	//讀取數據字串
	//"SELECT col from tablename;"
	Read_str := fmt.Sprintf("SELECT * FROM %s", TableName)

	//讀取
	//查詢多條
	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	rows, err := db.Query(Read_str)
	//defer 關閉查詢
	//一定要關閉(延遲)
	defer rows.Close()
	//檢查錯誤
	if err != nil {
		log.Fatal(err)
		return datas, err
	}

	//處理每一行
	//Next method 迭代查詢資料，回傳bool
	//func (rs *Rows) Next() bool
	for rows.Next() {
		var colname ColName
		//Scan method方法用來讀取每一列的值
		//func (rs *Rows) Scan(dest ...interface{}) error
		if err := rows.Scan(&colname.Name, &colname.Area, &colname.Total); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, colname)
	}

	//在迴圈中是否有錯誤
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return datas, err
	}
	return datas, nil
}

//查詢單一條件
func (db DB) ReadOne(TableName string, data ColName, col string, value string) (ColName, error) {
	//讀取數據字串
	//"SELECT col from tablename where ;"
	Read_str := fmt.Sprintf("SELECT * FROM %s where %s=%s", TableName, col, value)

	//讀取
	//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	row := db.QueryRow(Read_str)
	//defer 關閉查詢
	//一定要關閉(延遲)
	if err := row.Scan(&data.Name, &data.Area, &data.Total); err != nil {
		log.Fatal(err)
		return data, err
	}
	return data, nil
}
