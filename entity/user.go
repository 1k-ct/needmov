package entity

import (
	"time"
)

// Users users database -> id createdat updateat deletedat name password email
type Users struct {
	ID        int    `json:"id,omitempty"`
	CreatedAt string `json:"create_at,omitempty"`
	UpDatedAt string `json:"up_dated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name      string `json:"name,omitempty"`
	PassWord  string `json:"pass_word,omitempty"`
	Email     string `json:"email,omitempty"`
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
	VideoID   string `json:"video_id,omitempty"`
	VideoName string `json:"video_name,omitempty"`
	//VideoDescription string    `gorm:"type:text" json:"video_description,omitempty"`
	ThumbnailURL string    `json:"thumbnail_url,omitempty"`
	ViewCount    uint64    `gorm:"type:int" json:"video_count,omitempty"`
	CommentCount uint64    `gorm:"type:int" json:"comment_count,omitempty"`
	LikeCount    uint64    `gorm:"type:int" json:"like_count,omitempty"`
	DislikeCount uint64    `gorm:"type:int" json:"dislike_count,omitempty"`
	UploadDate   time.Time `json:"upload_date,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// ChannelInfos channel info
type ChannelInfos struct {
	//ID              uint64
	ChannelID       string    `json:"channel_id"`
	ChannelName     string    `json:"channel_name"`
	ViewCount       uint64    `gorm:"type:int" json:"view_count"`
	SubscriberCount uint64    `gorm:"type:int" json:"subscriber_count"`
	VideoCount      uint64    `gorm:"type:int" json:"video_count"`
	CreatedAt       time.Time `json:"create_at"`
}

/*
// ShiromiyaVideoInfos 白宮みみ
type ShiromiyaVideoInfos struct {
	//ID               int
	VideoID   string `json:"video_id,omitempty"`
	VideoName string `json:"video_name,omitempty"`
	//VideoDescription string    `gorm:"type:text" json:"video_description,omitempty"`
	ThumbnailURL string    `json:"thumbnail_url,omitempty"`
	ViewCount    uint64    `gorm:"type:int" json:"video_count,omitempty"`
	CommentCount uint64    `gorm:"type:int" json:"comment_count,omitempty"`
	LikeCount    uint64    `gorm:"type:int" json:"like_count,omitempty"`
	DislikeCount uint64    `gorm:"type:int" json:"dislike_count,omitempty"`
	UploadDate   time.Time `json:"upload_date,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// ShiromiyaChannelInfos 白宮みみ
type ShiromiyaChannelInfos struct {
	//ID              uint64
	ChannelID       string
	ChannelName     string
	ViewCount       uint64 `gorm:"type:int"`
	SubscriberCount uint64 `gorm:"type:int"`
	VideoCount      uint64 `gorm:"type:int"`
	CreatedAt       time.Time
}

// HashibaVideoInfos 羽柴なつみ
type HashibaVideoInfos struct {
	//ID               int
	VideoID   string `json:"video_id,omitempty"`
	VideoName string `json:"video_name,omitempty"`
	//VideoDescription string    `gorm:"type:text" json:"video_description,omitempty"`
	ThumbnailURL string    `json:"thumbnail_url,omitempty"`
	ViewCount    uint64    `gorm:"type:int" json:"video_count,omitempty"`
	CommentCount uint64    `gorm:"type:int" json:"comment_count,omitempty"`
	LikeCount    uint64    `gorm:"type:int" json:"like_count,omitempty"`
	DislikeCount uint64    `gorm:"type:int" json:"dislike_count,omitempty"`
	UploadDate   time.Time `json:"upload_date,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// HashibaChannelInfos 羽柴なつみ
type HashibaChannelInfos struct {
	//ID              uint64
	ChannelID       string
	ChannelName     string
	ViewCount       uint64 `gorm:"type:int"`
	SubscriberCount uint64 `gorm:"type:int"`
	VideoCount      uint64 `gorm:"type:int"`
	CreatedAt       time.Time
}
*/
// RegChannel データベースに登録するURLを保存
type RegChannel struct {
	ChannelID string `json:"channel_id"`
}
