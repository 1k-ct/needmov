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
	VideoID      string    `json:"video_id,omitempty"`
	VideoName    string    `json:"video_name,omitempty"`
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
	ChannelID       string    `json:"channel_id"`
	ChannelName     string    `json:"channel_name"`
	ViewCount       uint64    `gorm:"type:int" json:"view_count"`
	SubscriberCount uint64    `gorm:"type:int" json:"subscriber_count"`
	VideoCount      uint64    `gorm:"type:int" json:"video_count"`
	CreatedAt       time.Time `json:"create_at"`
}

// RegChannel データベースに登録するURLを保存
type RegChannel struct {
	ChannelID string `json:"channel_id"`
}

// Data コメントのデータ(api)
type Data struct {
	MasterChannelID string `json:"master_channel_id"`
	VideoURL        string `json:"video_url"`
	BadgeURL        string `json:"badge_url"`
	AuthorType      string `json:"author_type"`
	IsVerified      bool   `json:"is_verified"`
	IsChatOwner     bool   `json:"is_chat_owner"`
	IsChatSponsor   bool   `json:"is_chat_sponsor"`
	IsChatModerator bool   `json:"is_chat_moderator"`
	ChannelID       string `json:"channel_id"`
	Name            string `json:"name"`
	ImageURL        string `json:"image_url"`

	Type         string  `json:"type"`
	SID          string  `json:"sid"`
	Timestamp    int64   `gorm:"type:BIGINT" json:"timestamp"`
	ElapsedTime  string  `json:"elapsed_time"`
	Datetime     string  `json:"datetime"`
	Message      string  `gorm:"type:text" json:"message"`
	AmountValue  float64 `json:"amount_value"`
	AmountString string  `json:"amount_string"`
	Currency     string  `json:"currency"`
	BgColor      float64 `json:"bg_color"`
}
