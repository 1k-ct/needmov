package db

import (
	"log"
	"needmov/crypto"
	"needmov/entity"
	"os"

	_ "github.com/go-sql-driver/mysql" // gorm mysql import
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"google.golang.org/appengine"
)

//var (
//	db  *gorm.DB
//	err error
//)

// NewMakeDB dbの初期化　AutoMigrate dbの作成
func NewMakeDB() {
	db := ConnectGorm()
	defer db.Close()

	db.AutoMigrate(&entity.UsersMig{})
	db.AutoMigrate(&entity.Users{})
	db.AutoMigrate(&entity.ChannelInfos{}, &entity.VideoInfos{})
}

// CreateUser ユーザー登録
func CreateUser(username string, password string) []error {
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	// Encrypt 暗号化
	db := ConnectGorm()
	defer db.Close()
	// Insert処理
	if err := db.Create(&entity.UsersMig{Username: username, Password: passwordEncrypt}).GetErrors(); err != nil {
		return err
	}
	return nil
}

/*
func ConnectGorm() *gorm.DB { // localhost
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "sample"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}
*/
//ConnectGorm connect dbの接続
/*
func ConnectGorm() *gorm.DB { // 下のところは自分のものに変更してください
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)" //localhost
	DBNAME := "sample"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}
*/

//ConnectGorm connect dbの接続
func ConnectGorm() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	CONNECTIONNAME := os.Getenv("DB_CONNECTIONNAME")
	DBNAME := os.Getenv("DB_NAME")
	localConnection := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	cloudSQLConnection := USER + ":" + PASS + "@unix(/cloudsql/" + CONNECTIONNAME + ")/" + DBNAME + "?parseTime=true"
	var db *gorm.DB

	if appengine.IsAppEngine() {
		db, err = gorm.Open("mysql", cloudSQLConnection)
	} else {
		db, err = gorm.Open("mysql", localConnection)
	}
	if err != nil {
		panic(err.Error())
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
func GetDBContents() []entity.UsersMig {
	db := ConnectGorm()
	var users []entity.UsersMig
	db.Find(&users)
	db.Close()
	return users
}

// DeleteDB 選択したidをDBから削除
func DeleteDB(id int) {
	db := ConnectGorm()
	var user entity.UsersMig
	db.First(&user, id)
	db.Delete(&user)
	db.Close()
}

// GetUser ユーザーを一件取得
func GetUser(username string) entity.UsersMig {
	db := ConnectGorm()
	var user entity.UsersMig
	db.First(&user, "username = ?", username)
	db.Close()
	return user
}
