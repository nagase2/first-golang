package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)


func main() {
	//・・fmt.Println("🌟", aaa)
	header := [][]string{
		//ヘッダ行
		{"first_name", "last_name", "username"},
	}
	records := [][]string{
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	var filePath = "./dat2.log"

	isFileExist := true
	// ファイルが既にあるかどうかをチェック
	if file, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// file does not exist
		fmt.Println("ファイルがありません")
		isFileExist = false
	} else {
		// file exists
		fmt.Println("今の月は", int(time.Now().Month()))
		fmt.Println("ファイルの更新日時は・・・", int(file.ModTime().Month()))

		// ファイルの作成日時をチェックする
		fmt.Println("ファイルが作成されたのは月です。")
		fmt.Println("ファイルが更新されたのは月です")
	}

	// ファイルが存在する場合は追記、なければ新たに作成してファイルを開く
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	check(err)

	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(f)

	//　ファイルが存在しなかった場合はヘッダ行を出力
	if !isFileExist {
		if err := w.Write(header[0]); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	//----

	// //d2 := []byte{115, 111, 109, 101, 10}

	// for _, record := range records {

	// 	if err := f.Write(record); err != nil {
	// 		log.Fatalln("error writing record to csv:", err)
	// 	}
	// }
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n2)

	// It’s idiomatic to defer a Close immediately after opening a file.
	f.Sync()
	defer f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
