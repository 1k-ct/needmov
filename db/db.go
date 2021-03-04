package db

import (
	"errors"
	"needmov/crypto"
	"needmov/entity"

	_ "github.com/go-sql-driver/mysql" // gormでdbと接続は、mysqlでする。
	"github.com/jinzhu/gorm"
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
	db.AutoMigrate(&entity.RegChannel{})
	db.AutoMigrate(&entity.Data{})
}

// DropDBGorm []interface{}スライスで削除できる
func DropDBGorm(obj ...interface{}) {
	db := ConnectGorm()
	defer db.Close()
	db.DropTableIfExists(obj...)

}

// CreateUser ユーザー登録
func CreateUser(username string, password string) []error {
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	db := ConnectGorm()
	defer db.Close()
	if err := db.Create(&entity.UsersMig{Username: username, Password: passwordEncrypt}).GetErrors(); err != nil {
		return err
	}
	return nil
}

// ConnectGorm localhostの接続
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

/*

//ConnectGorm connect dbの接続 docker
func ConnectGorm() *gorm.DB {
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

//ConnectGorm connect dbの接続 本場
// func ConnectGorm() *gorm.DB {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	USER := os.Getenv("DB_USER")
// 	PASS := os.Getenv("DB_PASS")
// 	CONNECTIONNAME := os.Getenv("DB_CONNECTIONNAME")
// 	DBNAME := os.Getenv("DB_NAME")
// 	localConnection := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
// 	cloudSQLConnection := USER + ":" + PASS + "@unix(/cloudsql/" + CONNECTIONNAME + ")/" + DBNAME + "?parseTime=true"
// 	var db *gorm.DB

// 	if appengine.IsAppEngine() {
// 		db, err = gorm.Open("mysql", cloudSQLConnection)
// 	} else {
// 		db, err = gorm.Open("mysql", localConnection)
// 	}
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// AddNewInDB DBに新しく追加する
func AddNewInDB(id int, name string, password string, email string) {
	db := ConnectGorm()
	db.Create(&entity.Users{ID: id, Name: name, PassWord: password, Email: email})
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

// InsterRegChannel チャンネルURLをDBに登録
func InsterRegChannel(url string) ([]entity.RegChannel, error) {
	db := ConnectGorm()
	defer db.Close()
	regch, err := existenceDBValues(url) // urlのdb有無チェック！
	switch err {
	case true:
		db.Create(&entity.RegChannel{ChannelID: url})
		return regch, nil
	case false:
		return regch, errors.New("登録済みです")
	default:
		return regch, errors.New("おかしい")
	}
}

// existenceDBValues url (channel_id)がdbにあるか確認する　ありture なしfalse
func existenceDBValues(url string) ([]entity.RegChannel, bool) {
	db := ConnectGorm()
	defer db.Close()
	var regch []entity.RegChannel
	db.Where("channel_id = ?", url).Find(&regch)
	if len(regch) == 0 { //データベースに値が０個のときは成功、エラーではない
		return regch, true
	}
	return regch, false //errors.New("その、channel url はすでにあります")
}

// AllGetRegCh RegChannel(db)からすべての値を取る
func AllGetRegCh() []entity.RegChannel {
	db := ConnectGorm()
	defer db.Close()
	var regch []entity.RegChannel
	db.Find(&regch)
	return regch
}
