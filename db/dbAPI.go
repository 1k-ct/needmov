package db

// CreateDBSelfFn DB Insert create 自作関数
// var d Data
// db := db.ConnectGorm()
// db.Create(&d) 同じこと
func CreateDBSelfFn(d interface{}) error {
	db := ConnectGorm()
	defer db.Close()
	err := db.Create(d).Error
	if err != nil {
		return err
	}
	return nil
}
