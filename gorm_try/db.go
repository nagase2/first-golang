package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	fmt.Println("🌟initです！🌟")

	// mysql
	dsn := "docker:docker@tcp(127.0.0.1:9306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(mysql_db:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Globally
		//Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	RunMigrations()
	//データを再投入する
	Seed()
	DB.Logger = DB.Logger.LogMode(logger.Warn)
}

// 初期データを登録する
func Seed() {
	fmt.Println("🐮Seedの登録を行います。🐮")
	DB.Create(&Language{Code: "JP", Name: "日本語"})
	DB.Create(&Language{Code: "EN", Name: "英語"})
	DB.Create(&Blog{Title: "今日の話題1", Author: Author{Name: "山田　太郎"}})
	DB.Create(&Blog{Title: "今日の話題2", Author: Author{Name: "佐藤　太郎"}})
	DB.Create(&User{Name: "test", Age: 111,
		Friends: []*User{
			{Name: "友達1", Age: 99},
			{Name: "友達２", Age: 23},
			{Name: "友達３", Age: 95}}})
	DB.Create(&User{Name: "ナガセ", Age: 11,
		Friends: []*User{
			{Name: "友達1", Age: 99},
			{Name: "友達２", Age: 23},
			{Name: "友達３", Age: 95}}})

}

func RunMigrations() {

	fmt.Println("テーブルの再作成を行います。（データも全てリセット）")
	var err error
	allModels := []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}, &CreditCard{},
		&Product{}, &Model{}, &Blog{}, &Author{},
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	// テーブルを全てDropする
	DB.Migrator().DropTable("user_friends", "user_speaks")

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = DB.AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	// テーブルが全て作成されたかどうかをチェック
	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
}
