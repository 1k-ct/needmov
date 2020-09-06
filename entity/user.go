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

// VideoInfo video info
type VideoInfo struct {
	ID               int
	videoID          string
	videoName        string
	videoDescription string
	thumbnailURL     string
	viewCount        uint64
	commentCount     uint64
	likeCount        uint64
	dislikeCount     uint64
	uploadDate       time.Time
	createdAt        time.Time
}

// ChannelInfo channel info
type ChannelInfo struct {
	ID              uint64
	channelID       string
	channelName     string
	viewCount       uint64
	subscriberCount uint64
	videoCount      uint64
	createdAt       time.Time
}
