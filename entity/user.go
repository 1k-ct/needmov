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
	VideoDescription string
	ThumbnailURL     string
	ViewCount        uint64
	CommentCount     uint64
	LikeCount        uint64
	DislikeCount     uint64
	UploadDate       time.Time
	CreatedAt        time.Time
}

// ChannelInfos channel info
type ChannelInfos struct {
	//ID              uint64
	ChannelID       string
	ChannelName     string
	ViewCount       uint64
	SubscriberCount uint64
	VideoCount      uint64
	CreatedAt       time.Time
}
