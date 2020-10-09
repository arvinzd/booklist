package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/errors/fmt"
)

var db *sqlx.DB

type Book struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Price int    `db:"price"`
}

func Init(addr, user, passwd, dbname string) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, passwd, addr, dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		return
	}
	// 最大连接
	db.SetMaxOpenConns(100)
	// 最大空闲
	db.SetMaxIdleConns(16)
	return
}

func QueryBook() (booklist []*Book, err error) {
	sqlStr := "select * from book"
	err = db.Select(&booklist, sqlStr)
	if err != nil {
		return nil, err
	}
	return booklist, nil
}

func InsertBook(title string, price int) (err error) {
	sqlStr := "insert into book (title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		return err
	}
	return
}

func DeleteBook(id int) (err error) {
	sqlStr := "delete from book where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}
