package db

import (
	"database/sql"
	"fmt"
	"log"
	//キャスト用
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// DBからユーザーの単語と意味を取得する関数
func Word_display(id int) ([]string, []string) {
	//sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", "")

	if err != nil {
		fmt.Println("SQL接続失敗")
		panic(err.Error())
	}

	defer db.Close()
	string_id := strconv.Itoa(id) //int→stringにキャスト
	rows, err := db.Query("SELECT word,meaning FROM word WHERE id=\"" + string_id + "\";")
	if err != nil {
		fmt.Println("select失敗")
		log.Fatal(err)
	}
	var (
		word    *string
		meaning *string
	)

	var word_list, meaning_list []string
	for rows.Next() {
		if err := rows.Scan(&word, &meaning); err != nil {
			fmt.Println("スキャンに失敗")
			log.Fatal(err)
		}
		word_list = append(word_list, *word)
		meaning_list = append(meaning_list, *meaning)
	}
	return word_list, meaning_list
}
