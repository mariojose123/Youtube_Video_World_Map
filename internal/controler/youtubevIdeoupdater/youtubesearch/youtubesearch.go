package YoutubeAPISearchOptions

import (
	"context"
	. "guess_music_youtube/internal/auxfunctions"
	"os"
	"strconv"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeAPISearchOptions struct {
	RadiusUnit      string
	safeSearch      string
	order           string
	maxresults      int64
	isSyndicated    string
	PublishedBefore string
	RegionCodeISO   string
	TopicId         string
	Radius          string
}

func (searchoptions *YoutubeAPISearchOptions) SetEnviromentVariablesSearch() error {
	var err error

	searchoptions.RadiusUnit = os.Getenv("SEARCH_VID_RADIUSUNIT")
	searchoptions.Radius = os.Getenv("SEARCH_VID_RADIUS")

	searchoptions.safeSearch = os.Getenv("SEARCH_VID_SAFESEARCH")

	searchoptions.order = os.Getenv("SEARCH_VID_ORDER")
	searchoptions.isSyndicated = os.Getenv("SEARCH_VID_ISSYNDICATED")
	searchoptions.TopicId = os.Getenv("SEARCH_VID_TOPICID")
	searchoptions.PublishedBefore = os.Getenv("SEARCH_VID_PUBLISHEDBEFORE")

	searchoptions.maxresults, err = strconv.ParseInt(os.Getenv("SEARCH_VID_MAXRESULTS"), 10, 64)
	if err != nil {
		return ReturnError(err)
	}
	return nil
}
func (searchoptions YoutubeAPISearchOptions) CallYoutubeAPISearch(ISO2 string, lat string, long string) ([]*youtube.SearchResult, error) {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		return nil, ReturnError(err)
	}

	call := service.Search.List([]string{"snippet"})
	response, err := call.RegionCode(ISO2).
		VideoEmbeddable("true").
		Location(lat + "," + long).
		LocationRadius(searchoptions.Radius + searchoptions.RadiusUnit).
		VideoSyndicated(searchoptions.isSyndicated).
		TopicId(searchoptions.TopicId).
		Type("video").
		SafeSearch(searchoptions.safeSearch).
		Order(searchoptions.order).
		PublishedBefore(searchoptions.PublishedBefore).
		MaxResults(searchoptions.maxresults).Do()
	if err != nil {
		return nil, ReturnError(err)
	}
	return response.Items, nil
}
