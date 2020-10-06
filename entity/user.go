package entity

import (
	"time"
)

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

// VideoInfos video info
type VideoInfos struct {
	//ID               int
	VideoID          string
	VideoName        string
	VideoDescription string `gorm:"type:text"`
	ThumbnailURL     string
	ViewCount        uint64 `gorm:"type:int"`
	CommentCount     uint64 `gorm:"type:int"`
	LikeCount        uint64 `gorm:"type:int"`
	DislikeCount     uint64 `gorm:"type:int"`
	UploadDate       time.Time
	CreatedAt        time.Time
}

// ChannelInfos channel info
type ChannelInfos struct {
	//ID              uint64
	ChannelID       string
	ChannelName     string
	ViewCount       uint64 `gorm:"type:int"`
	SubscriberCount uint64 `gorm:"type:int"`
	VideoCount      uint64 `gorm:"type:int"`
	CreatedAt       time.Time
}
