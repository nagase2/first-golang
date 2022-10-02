package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// // equals
// type Blog struct {
// 	ID      int64
// 	Name    string
// 	Email   string
// 	Upvotes int32
// }

func main() {
	fmt.Println("Hello, world.2")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			//LogLevel:                  logger.Info, // Log level
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	// // sql lite
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	//
	// DB.Model(&Product{}).Create(map[string]interface{}{
	// 	"Codec": "jinzhu", "Price": 18,
	// })

	cp := Product{Code: "D41", Price: 100}
	// Create
	DB.Create(&cp)
	fmt.Printf("🈲🈲🈲")

	DB.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//DB.First(&product, 1)                 // find product with integer primary key
	DB.First(&product, "code = ?", "D42") // find product with code D42

	// Update
	DB.Model(&product).Update("Price", 88200)

	//-----------------------

	// Create index for Name field
	DB.Migrator().CreateIndex(&User{}, "Name")
	DB.Migrator().CreateIndex(&User{}, "idx_name")

	DB.Create(&User{
		Name: "testuser1",
		//CreditCards: [&CreditCard{Number: "411111111111"}]
	})
	DB.Create(&User{
		Name: "22222",
		//CreditCards: [&CreditCard{Number: "411111111111"}]
	})

	// Continuous session mode
	tx := DB.Session(&gorm.Session{Logger: newLogger})
	tx.First(&User{}, 1)

	var users []User
	DB.Find(&users)

	// var product2 Product
	// tx.First(&product2{}, 1)
	// DB.Model(&product2{Code: "42"}).Update("Price", 11200)

	type Result struct {
		ID   int
		Name string
		Age  int
	}
	// raw SQLで検索
	var results []Result
	//DB.Raw("SELECT id, name, age FROM users WHERE age > ?", 20).Scan(&results)
	DB.Raw("SELECT * FROM users WHERE (name = @name AND age > @age)",
		map[string]interface{}{"name": "test", "age": 11}).Find(&results)

	for index, result := range results {
		fmt.Println("🐵", index, result.ID, result.Name)
	}

	// 取得したデータを全て出力
	// for index, user := range users {
	// 	fmt.Println(index, user.Name)
	// }
}
