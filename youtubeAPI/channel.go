package youtubeapi

import (
	"errors"
	"log"
	"needmov/entity"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// PrintChannelInfo return (*entity.ChannelInfos, error)
func PrintChannelInfo(channelID string) (*entity.ChannelInfos, error) {
	service := newYoutubeService(newClient())
	call := service.Channels.List([]string{"snippet", "contentDetails", "statistics"}).Id(channelID).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return nil, errors.New("response error")
	}
	if response.Items == nil {
		return nil, errors.New("IDが無効です")
	}
	item := response.Items[0]
	c := &entity.ChannelInfos{
		ChannelID:       item.Id,
		ChannelName:     item.Snippet.Title,
		ViewCount:       item.Statistics.ViewCount,
		SubscriberCount: item.Statistics.SubscriberCount,
		VideoCount:      item.Statistics.VideoCount,
	}
	return c, nil
}
func newClient() *http.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	API_KEY := os.Getenv("API_KEY")
	client := &http.Client{
		Transport: &transport.APIKey{Key: (API_KEY)}, // ここ、API KEY
	}
	return client
}

func newYoutubeService(client *http.Client) *youtube.Service {
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}
	return service
}

// PrintVideoInfo return(*entity.VideoInfos, error)
func PrintVideoInfo(videoID string) (*entity.VideoInfos, error) {
	service := newYoutubeService(newClient())
	call := service.Videos.List([]string{"id,snippet,Statistics"}).Id(videoID).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return nil, errors.New("respone error")
	}
	item := response.Items[0]
	uploadDate, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
	if err != nil {
		return nil, errors.New("uploadDate error")
	}
	v := &entity.VideoInfos{
		VideoID:      item.Id,
		VideoName:    item.Snippet.Title,
		ThumbnailURL: item.Snippet.Thumbnails.High.Url,
		ViewCount:    item.Statistics.ViewCount,
		CommentCount: item.Statistics.CommentCount,
		LikeCount:    item.Statistics.LikeCount,
		DislikeCount: item.Statistics.DislikeCount,
		UploadDate:   uploadDate,
	}
	return v, nil
}
