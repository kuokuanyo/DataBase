//刪除資料庫表
package conn

import (
	"database/sql"
	"fmt"
)

//funciton(delete database)
func Delete_Db(db *sql.DB, DbName string) {

	//刪除字串
	Delete := fmt.Sprintf("DROP DATABASE %s", DbName)

	//刪除
	db.Exec(Delete)
}

//function(delete table)
func Delete_Tb(db *sql.DB, TableName string) {

	//刪除字串
	Delete := fmt.Sprintf("DROP TABLE %s", TableName)

	//刪除
	db.Exec(Delete)
}
