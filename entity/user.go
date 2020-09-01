package entity

// Users users database -> id createdat updateat deletedat name password email
type Users struct {
	ID        int
	CreatedAt string
	UpDatedAt string
	DeletedAt string
	Name      string
	PassWord  string
	Email     string
}

// UsersMig -> gorm.Model UserName Password
type UsersMig struct {
	ID       uint   `form:"id" gorm:"primaryKey"`
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

// SessionInfo UserID type is interface{}
type SessionInfo struct {
	ID interface{}
}
