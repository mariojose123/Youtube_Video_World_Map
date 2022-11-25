package Webapp

import (
	"fmt"
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

func (control Controler) GetCountryVideos(numberofvideos int, country string) ([]*youtubevideo.Youtubevideo, error) {
	numberofvideosstring := strconv.Itoa(numberofvideos)
	rows, err := control.db.QueryRowsVideo(country, numberofvideosstring)
	if err != nil {
		fmt.Println("db.QueryVideos", err)
	}
	var youtubevideos []*youtubevideo.Youtubevideo
	for rows.Next() {
		c := new(youtubevideo.Youtubevideo)
		err := rows.Scan(
			&c.VideoID,
			&c.Iso2,
			&c.Lat,
			&c.Long,
			&c.Radius,
			&c.Tagid,
			&c.Description,
			&c.Cityid,
			&c.Country)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		youtubevideos = append(youtubevideos, c)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return youtubevideos, nil
}
