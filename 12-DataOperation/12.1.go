package main

import (
	"Study/MSZL/Study/12-DataOperation/pojo"
	_ "Study/MSZL/Study/12-DataOperation/pojo"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type User pojo.User

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_study")
	if err != nil {
		panic(err)
	}
	//最大空闲连接数，默认不配置 => 最大两个连接数
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置 => 不限制最大连接数
	db.SetMaxOpenConns(100)
	//连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)

	err = db.Ping()
	if err != nil {
		log.Println("数据库连接失败")
		db.Close()
		panic(err)
	}
	DB = db
}

func _insert() {
	r, err := DB.Exec("insert into user(username,sex,email) values(?,?,?)", "NanCheng", "男", "3344567")
	if err != nil {
		log.Println("执行Sql语句失败")
		panic(err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("插入成功", id)
}

func _select(id int) (*User, error) {
	rows, err := DB.Query("select * from user where user_id = ?", id)
	if err != nil {
		log.Println("查询数据库出现错误:", err)
		return nil, errors.New(err.Error())
	}
	user := new(User)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.Username, &user.Sex, &user.Email); err != nil {
			log.Println("扫描数据出现错误:", err)
			return nil, errors.New(err.Error())
		}
	}
	return user, nil
}

func _update(username string, id int) {
	ret, err := DB.Exec("update user set username=? where user_id=?", username, id)
	if err != nil {
		log.Println("更新出现问题:", err)
		return
	}
	affected, _ := ret.RowsAffected()
	fmt.Println("更新成功的行数:", affected)
}

func _delete(id int) {
	ret, err := DB.Exec("delete from user where user_id=?", id)
	if err != nil {
		log.Println("删除出现问题:", err)
		return
	}
	affected, _ := ret.RowsAffected()
	fmt.Println("删除成功的行数:", affected)
}

func main() {
	defer DB.Close()
	fmt.Println(_select(1))
}

func insertTx(username string) {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Println("开启事务错误:", err)
		return
	}
	// 执行语句
	ret, err := tx.Exec("insert into user (username,sex,email) values (?,?,?)", username, "man", "test@mszlu.com")
	if err != nil {
		log.Println("事务sql执行出错:", err)
		return
	}
	id, _ := ret.LastInsertId()
	fmt.Println("插入成功:", id)
	// 回滚操作
	if username == "lisi" {
		fmt.Println("回滚...")
		_ = tx.Rollback()
	} else {
		// 提交操作
		_ = tx.Commit()
	}

}
