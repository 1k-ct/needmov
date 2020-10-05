package youtubeapi

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// PrintChannelInfo return id string, name string, viewCount uint64, subscriberCount uint64, videoCount uint64,
func PrintChannelInfo(channelID string) (string, string, uint64, uint64, uint64) {
	service := newYoutubeService(newClient())
	//lis := []string{"snippet", "contentDetails", "statistics"}
	call := service.Channels.List([]string{"snippet", "contentDetails", "statistics"}).
		Id(channelID).
		MaxResults(1)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	item := response.Items[0]

	id := item.Id
	name := item.Snippet.Title
	//description := item.Snippet.Description
	//thumbnailURL := item.Snippet.Thumbnails.High.Url
	//playlistID := item.ContentDetails.RelatedPlaylists.Uploads
	viewCount := item.Statistics.ViewCount
	subscriberCount := item.Statistics.SubscriberCount
	videoCount := item.Statistics.VideoCount
	/*
		fmt.Printf("channel id: %v\n\nチャンネル名: \n%v\n\n説明: %v\n\nサムネイルURL: %v\n\nplaylist id: %v\n\n総再生回数: %v\n\nチャンネル登録者数: %v\n\n動画数: %v\n",
			id,
			name,
			description,
			thumbnailURL,
			playlistID,
			viewCount,
			subscriberCount,
			videoCount,
		)
	*/
	return id, name, viewCount, subscriberCount, videoCount
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

// PrintVideoInfo return(id string, name string, description string, thumbnailURL string, viewCount uint64, commentCount uint64, likeCount uint64, dislikeCount uint64, uploadData time.Time)
func PrintVideoInfo(videoID string) (string, string, string, string, uint64, uint64, uint64, uint64, time.Time) {
	service := newYoutubeService(newClient())
	call := service.Videos.List([]string{"id,snippet,Statistics"}).
		Id(videoID).
		MaxResults(1)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	item := response.Items[0]
	id := item.Id
	name := item.Snippet.Title
	description := item.Snippet.Description
	thumbnailURL := item.Snippet.Thumbnails.High.Url
	viewCount := item.Statistics.ViewCount
	commentCount := item.Statistics.CommentCount
	likeCount := item.Statistics.LikeCount
	dislikeCount := item.Statistics.DislikeCount
	//channelID := item.Snippet.ChannelId
	//categoryID := item.Snippet.CategoryId
	//categoryName := getVideoCategory(categoryID)
	uploadDate, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return id, name, description, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate
}
