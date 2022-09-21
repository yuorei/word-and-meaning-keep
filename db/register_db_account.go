package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ユーザー登録の関数
func User_db(name string, password string) {
	//sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", "")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	rows, err := db.Query("SELECT id FROM user WHERE id = (SELECT MAX(id) FROM user);")
	if err != nil {
		log.Fatal(err)
	}
	var id *int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			fmt.Println("idスキャンに失敗")
			log.Fatal(err)
		}
	}
	// プリペアードステートメントを使用
	in, err := db.Prepare("INSERT INTO user(id ,name,password) VALUES(?,?,?)")
	if err != nil {
		fmt.Println("データベース接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("データベース接続成功")
	}
	defer db.Close() //最後に閉じる
	//insert
	result, err := in.Exec(*id+1, name, password)
	if err != nil {
		panic(err.Error())
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastId)
}
