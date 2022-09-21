package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 単語をDBに登録する関数
func Word_register(id int, word string, meaning string) {
	//sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", "")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// プリペアードステートメントを使用
	in, err := db.Prepare("INSERT INTO word(id ,word,meaning) VALUES(?,?,?)")

	if err != nil {
		fmt.Println("データベース接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("データベース接続成功")
	}
	defer db.Close() //最後に閉じる
	//insert
	result, err := in.Exec(id, word, meaning) //idはユーザーアカウント
	if err != nil {
		fmt.Println("insert失敗")
		panic(err.Error())
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastId)
}
