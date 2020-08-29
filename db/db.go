package db

import (
	"fmt"

	"needmov/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// OpenDB is
func OpenDB() *gorm.DB {
	DBMS := "mysql"
	USER := "user1"
	PASS := "Password_01"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "users"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	fmt.Println("db connected: ", &db)
	return db
}

//ConnectGorm connect db
func ConnectGorm() *gorm.DB { // 下のところは自分のものに変更してください
	DBMS := "mysql"
	USER := "user1"
	PASS := "Password_01"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_sample"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}

// AddNewInDB DBに新しく追加する
func AddNewInDB(id int, name string, password string, email string) { //, createdAt string
	db := ConnectGorm()
	db.Create(&entity.Users{ID: id, Name: name, PassWord: password, Email: email}) //, CreatedAt: createdAt
	defer db.Close()
}

// GetDBContents DBの全ての投稿を取得する
func GetDBContents() []entity.Users {
	db := ConnectGorm()
	var users []entity.Users
	db.Find(&users)
	db.Close()
	return users
}

// DeleteDB 選択したidをDBから削除
func DeleteDB(id int) {
	db := ConnectGorm()
	var user entity.Users
	db.First(&user, id)
	db.Delete(&user)
	db.Close()
}
