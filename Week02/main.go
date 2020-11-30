package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"test/dao"
)

func main()  {
	db, err := sql.Open("mysql", "root:111@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return
	}
	result, err := dao.QueryRow(db, "select result from tableInfo where name=?", "lilw")
	if dao.IsQueryEmpty(errors.Cause(err)) {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println(result)
}
