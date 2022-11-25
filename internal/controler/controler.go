package controler

import (
	"database/sql"
	"fmt"
	. "guess_music_youtube/internal/auxfunctions"
	Database "guess_music_youtube/internal/database"
	"guess_music_youtube/internal/youtubevideo"
	"strconv"
)

type Controler struct {
	db Database.Database
}

func (control Controler) NewControler(db Database.Database) Controler {
	return Controler{db}
}

//Get a random from every Country
func (control Controler) GetRandomEveryCountry() ([]*youtubevideo.Youtubevideo, error) {
	rows, err := control.db.QueryRandomVideoFromEveryCountry()
	if err != nil {
		return nil, ReturnError(err)
	}
	youtubevideos, err := SQLRowstoYoutubeVideoStructs(rows)
	if err != nil {
		return nil, ReturnError(err)
	}
	return youtubevideos, nil
}

func (control Controler) GetCountryVideos(numberofvideos int, country string) ([]*youtubevideo.Youtubevideo, error) {
	numberofvideosstring := strconv.Itoa(numberofvideos)
	rows, err := control.db.QueryRowsVideo(country, numberofvideosstring)
	if err != nil {
		return nil, ReturnError(err)
	}
	var youtubevideos []*youtubevideo.Youtubevideo
	youtubevideos, err = SQLRowstoYoutubeVideoStructs(rows)
	if err != nil {
		return nil, ReturnError(err)
	}
	return youtubevideos, nil
}

func SQLRowstoYoutubeVideoStructs(rows *sql.Rows) ([]*youtubevideo.Youtubevideo, error) {
	var youtubevideos []*youtubevideo.Youtubevideo
	for rows.Next() {
		c := new(youtubevideo.Youtubevideo)
		err := rows.Scan(
			&c.VideoID,
			&c.ChannelID,
			&c.Iso2,
			&c.Lat,
			&c.Long,
			&c.Radius,
			&c.Tagid,
			&c.Description,
			&c.Cityid,
			&c.Country,
			&c.ThumbnailUrl,
			&c.PublishedAt)
		if err != nil {
			fmt.Println(err)
			return nil, ReturnError(err)
		}

		youtubevideos = append(youtubevideos, c)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, ReturnError(err)
	}

	return youtubevideos, nil
}
