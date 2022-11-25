package updater

import (
	"fmt"
	. "guess_music_youtube/internal/auxfunctions"
	. "guess_music_youtube/internal/controler/city"
	. "guess_music_youtube/internal/controler/youtubevIdeoupdater/youtubesearch"
	. "guess_music_youtube/internal/database"
	"log"
	"time"
)

type VideoUpdater struct {
	db            Database
	iso2index     int
	videoPerDay   int
	days          int
	cities        []*City
	searchoptions YoutubeAPISearchOptions
}

//Update Database with youtube video of certain types according to env variables */
func (vu *VideoUpdater) UpdaterMain() error {

	err := vu.searchoptions.SetEnviromentVariablesSearch()
	if err != nil {
		return ReturnError(err)
	}
	vu.cities, err = vu.getaCityForEachCountry()
	if err != nil {
		return ReturnError(err)
	}
	err = vu.UpdateVideoWithCities()
	if err != nil {
		return ReturnError(err)
	}
	return nil
}

func (vu *VideoUpdater) Init(VideoPerDay int, Days int, db Database) {
	vu.videoPerDay = VideoPerDay
	vu.days = Days
	vu.db = db
}

/*Search for videos for every city in loop videoPerDay per days VideoUpdater atributes*/
func (vu VideoUpdater) UpdateVideoWithCities() error {
	cities := vu.cities
	log.Println(cities)
	days := vu.days
	videoperday := vu.videoPerDay
	for {
		for indexday := 0; indexday < days; indexday++ {
			for indexsearch := 0; indexsearch < videoperday; indexsearch++ {
				indexc := (indexday*videoperday + indexsearch) % len(cities)
				city := cities[indexc]
				youtubevideos, err := vu.searchoptions.CallYoutubeAPISearch(city.Iso2, city.Lat, city.Long)
				if err != nil {
					return ReturnError(err)
				}
				log.Print(city.Iso2 + "\n")
				for _, video := range youtubevideos {
					vu.db.InsertYoutubeVideos(
						video.Id.VideoId,
						video.Snippet.ChannelId,
						city.Iso2,
						city.Lat,
						city.Long,
						vu.searchoptions.Radius+vu.searchoptions.RadiusUnit,
						vu.searchoptions.TopicId,
						video.Snippet.Description,
						city.CityId,
						city.Country,
						video.Snippet.Thumbnails.Default.Url,
						video.Snippet.PublishedAt,
					)
				}
			}
			time.Sleep(time.Hour * 24)
		}
	}
}

//Get most populated city for each country
func (vu *VideoUpdater) getaCityForEachCountry() ([]*City, error) {
	rows, err := vu.db.QueryMostPopulatedCityPerCountry()
	if err != nil {
		return nil, ReturnError(err)
	}
	var cities []*City
	for rows.Next() {
		c := new(City)
		err := rows.Scan(&c.Iso2,
			&c.Lat,
			&c.Long,
			&c.CityId,
			&c.Country,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		cities = append(cities, c)
	}
	return cities, nil
}
