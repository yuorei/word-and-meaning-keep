package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ログインのための関数
func User_login(name string, password string) int {
	//sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", "")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	rows, err := db.Query("SELECT id,name,password FROM user WHERE name=\"" + name + "\"and password=\"" + password + "\";")
	if err != nil {
		fmt.Println("select失敗")
		log.Fatal(err)
	}
	var (
		id *int
	)

	for rows.Next() {
		if err := rows.Scan(&id, &name, &password); err != nil {
			fmt.Println("スキャンに失敗")
			log.Fatal(err)
		}
	}
	m := 0
	if id == nil {
		//何もない場合は入力が間違いとして戻り値を0で返す
		id = &m
		return *id
	} else {
		return *id
	}

}
