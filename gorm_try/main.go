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

	user1 := User{
		Name: "jinzhu",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
			{Name: "TK"},
		},
		Company: Company{Name: "北海道電力"},
	}
	user2 := User{
		Name: "田中　ひろたか",
		Languages: []Language{
			{Name: "JP"},
			{Name: "IN"},
		},
		Company: Company{Name: "中部電力"},
	}
	DB.Create(&user1)
	DB.Create(&user2)

	// UpSeart（なんで？）
	lang1 := Language{
		Code: "FN",
		Name: "フランス"}
	DB.Create(&lang1)
	lang2 := Language{
		Code: "FNA",
		Name: "AAA"}
	DB.Create(&lang2)
	fmt.Println("Language削除します")
	DB.Delete(&lang1)

	fmt.Println("ユーザ削除します")
	// DB.Delete(&user)
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

	// こうすれば受け取るかたちのStructを作らなくても結果を受け取れる
	mapResult := map[string]interface{}{}
	DB.Model(&User{}).First(&mapResult, "id = ?", 1)

	fmt.Println("😸", mapResult["name"], mapResult["age"])

	// JOINしてデータをとってくる
	type result1 struct {
		Name        string
		CompanyName string
	}
	DB.Model(&User{}).Select("users.name, companies.name").Joins("left join companies on  users.company_id=companies.id").Scan(&result1{})
	fmt.Println()
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

	// ORMの機能で、フランス語がはなせるUserを検索する

	// 取得したデータを全て出力
	// for index, user := range users {
	// 	fmt.Println(index, user.Name)
	// }
}


func (u *Language) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Print("🌟🌟🌟RoleCheck（ダミー）🌟🌟🌟", u)
	//   if u.Role == "admin" {
	//     return errors.New("admin user not allowed to delete")
	//   }
	return
}
