package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMySQL() (err error) {
	dsn := "root:1234qwer@tcp(127.0.0.1:60042)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(200)                 //最大连接数
	db.SetMaxIdleConns(10)                  //连接池里最大空闲连接数。必须要比maxOpenConns小
	db.SetConnMaxLifetime(time.Second * 10) //最大存活保持时间
	db.SetConnMaxIdleTime(time.Second * 10) //最大空闲保持时间
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
	} else {
		fmt.Println("connect to db success")
	}
	select_one()
	queryMultiRowDemo()
	insertRowDemo()
	updateRowDemo()
	deleteRowDemo()
	defer db.Close()

}

type user struct {
	id int
	// age  int
	name string
}

//查询
func select_one() {
	sqlStr := "SELECT id, name FROM sys_user WHERE id=?"
	var u user
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name)
	if err != nil {
		fmt.Printf("scan failed， err: %v\n", err)
		return
	}
	fmt.Printf("id:%d,name:%s\n", u.id, u.name)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "SELECT id,name FROM sys_user WHERE id>?"
	var u user
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	//非常重要，关闭rows释放持有的数据库链接
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id:%d,name:%s\n", u.id, u.name)
	}
}

//插入数据
func insertRowDemo() {
	sqlStr := "INSERT INTO sys_user(id,name) VALUES(?,?)"
	ret, err := db.Exec(sqlStr, 23, "javapub")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	var theID int64
	theID, err = ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theID)
}

//更新数据
func updateRowDemo() {
	sqlStr := "UPDATE sys_user SET name=? WHERE id=?"
	ret, err := db.Exec(sqlStr, "i love javapub", 23)
	if err != nil {
		fmt.Printf("updated failed,err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed,err:%v\b", err)
		return
	}
	fmt.Printf("updated success, the rows affected is %d\n", n)
}

//删除数据
func deleteRowDemo() {
	sqlStr := "DELETE FROM sys_user WHERE id=?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
