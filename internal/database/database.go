package Database

import (
	"context"
	"database/sql"
	"fmt"
	. "guess_music_youtube/internal/auxfunctions"
	"log"
	"net/url"
	"os"
	"strconv"
)

type Database struct {
	dsn           url.URL
	createtables  bool
	sslenabled    bool
	videoDBname   string
	channelDBname string
	CityDBname    string
}

func (db *Database) setEnviromentVariables() error {
	var err error
	db.createtables, err = strconv.ParseBool(os.Getenv("CREATE_TABLES"))
	if err != nil {
		return ReturnError(err)
	}
	db.sslenabled, err = strconv.ParseBool(os.Getenv("SLL_ENABLED"))
	if err != nil {
		return ReturnError(err)
	}
	scheme := os.Getenv("DB_SCHEME")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	path := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	db.CityDBname = os.Getenv("DB_CITYNAME")
	db.channelDBname = os.Getenv("DB_CHANNELNAME")
	db.videoDBname = os.Getenv("DB_VIDEONAME")

	db.dsn = url.URL{
		Scheme: scheme,
		Host:   host + ":" + port,
		User:   url.UserPassword(user, password),
		Path:   path,
	}
	return nil
}

//Start Database and Create Tables if CreateDataTables is true
func (db *Database) InitDb() error {
	err := db.setEnviromentVariables()
	if err != nil {
		return ReturnError(err)
	}
	q := db.dsn.Query()
	if !db.sslenabled {
		q.Add("sslmode", "disable")
	}
	db.dsn.RawQuery = q.Encode()

	if db.createtables {
		db.createTables()
	}
	return nil
}

//Create Tables for VideoGuessAplication
func (db *Database) createTables() error {
	err := db.WORLDPointsDB()
	if err != nil {
		log.Print("WORLDPointsDB")
		return ReturnError(err)
	}
	err = db.CreateYoutubeVideoDB()
	if err != nil {
		log.Print("CreateYoutubeVideoDB")
		return ReturnError(err)

	}
	err = db.CsvDatabaseCity("/data/worldcities.csv")
	if err != nil {
		log.Print("CsvDatabaseCity")
		return ReturnError(err)
	}
	return nil
}

func (db *Database) CreateYoutubeVideoDB() error {
	err := db.Exec(`CREATE TABLE ` + db.videoDBname + `(
	VIDEOID varchar NOT NULL,
	CHANNELID varchar NOT NULL,
	ISO2 VARCHAR,
	LAT DOUBLE PRECISION,
	LONG DOUBLE PRECISION,
	RADIUS VARCHAR,
	TAGID varchar,
	TITLE varchar,
	DESCRIPTION varchar,
	CITYID int,
	COUNTRY VARCHAR,
	SEARCHTERM varchar,
	THUMBNAIL varchar,
	PUBLISHEDAT varchar,
	PRIMARY KEY (VIDEOID),
	FOREIGN KEY (CITYID) REFERENCES ` + db.CityDBname + `(CITYID));`)
	if err != nil {
		return ReturnError(err)

	}
	return nil
}

func (db *Database) CreateChannelDB() error {
	err := db.Exec(`CREATE TABLE ` + db.channelDBname + `(
		CHANNELID varchar,
		TITLE     varchar,
		PRIMARY KEY (CHANNELID))`)
	if err != nil {
		return ReturnError(err)

	}
	return nil
}

/* Create Database World Points CITIES*/
func (db *Database) WORLDPointsDB() error {
	err := db.Exec(
		`CREATE TABLE  ` + db.CityDBname + `(
		CITY VARCHAR,
		CITYASCII VARCHAR,
		LATITUDE DOUBLE PRECISION,
		LONGITUDE DOUBLE PRECISION,
		COUNTRY VARCHAR,
		ISO2 CHAR(2),
		ISO3 CHAR(3),
		ADMINNAME VARCHAR,
		CAPITAL VARCHAR,
		POPULATION VARCHAR,
		CITYID INT NOT NULL,
		PRIMARY KEY (CITYID)
		);
  `)
	if err != nil {
		return ReturnError(err)

	}
	return nil
}

/* CSV of cities to insert into DATABASE db.CityDBname*/
func (db *Database) CsvDatabaseCity(url string) error {
	err := db.Exec(`COPY ` + db.CityDBname + `
	FROM '` + url + `'
	DELIMITER ','
	CSV HEADER;`)
	if err != nil {
		return ReturnError(err)

	}
	return nil
}

/*Add Videos in the Database  */
func (db *Database) InsertYoutubeVideos(videoID string, channelID string, ISO2 string,
	lat string, long string, radius string, tagid string,
	description string, cityid string, country string, url string, publishat string) error {
	opendb, err := sql.Open("pgx", db.dsn.String())
	if err != nil {
		return ReturnError(err)

	}
	defer func() {
		_ = opendb.Close()
		fmt.Println("closed")
	}()
	if err := opendb.PingContext(context.Background()); err != nil {
		return ReturnError(err)

	}
	_, err = opendb.Exec(
		`INSERT INTO VIDEOS (VIDEOID,CHANNELID,ISO2,LAT,LONG,RADIUS,TAGID,DESCRIPTION,CITYID,COUNTRY,THUMBNAIL,PUBLISHEDAT
)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, videoID,
		channelID,
		ISO2,
		lat,
		long,
		radius,
		tagid,
		"'"+description+"'",
		cityid,
		"'"+country+"'",
		"'"+url+"'",
		publishat)
	if err != nil {
		return ReturnError(err)
	}
	return nil
}

func (db *Database) Exec(query string) error {
	opendb, err := sql.Open("pgx", db.dsn.String())
	if err != nil {
		return ReturnError(err)

	}
	defer func() {
		_ = opendb.Close()
		fmt.Println("closed")
	}()
	if err := opendb.PingContext(context.Background()); err != nil {
		return ReturnError(err)

	}
	_, err = opendb.ExecContext(context.Background(), query)
	if err != nil {
		return ReturnError(err)

	}
	return nil
}

func (db *Database) QueryMostPopulatedCityPerCountry() (*sql.Rows, error) {
	opendb, err := sql.Open("pgx", db.dsn.String())
	if err != nil {
		return nil, ReturnError(err)

	}
	defer func() {
		_ = opendb.Close()
		fmt.Println("closed")
	}()
	if err := opendb.PingContext(context.Background()); err != nil {
		return nil, ReturnError(err)

	}
	rows, err := opendb.QueryContext(context.Background(),
		`SELECT DISTINCT ON(ISO2) ISO2 , LATITUDE , LONGITUDE , CITYID , COUNTRY 
	FROM `+db.CityDBname+` ORDER BY ISO2 , POPULATION DESC;`)
	if err != nil {
		return nil, ReturnError(err)
	}
	return rows, nil
}

//Select Video According to ISO2,I should develop more options to sql where
func (db *Database) QueryRowsVideo(ISO2 string, numberofvid string) (*sql.Rows, error) {
	opendb, err := sql.Open("pgx", db.dsn.String())
	if err != nil {
		return nil, ReturnError(err)

	}
	defer func() {
		_ = opendb.Close()
		fmt.Println("closed")
	}()
	if err := opendb.PingContext(context.Background()); err != nil {
		return nil, ReturnError(err)

	}
	rows, err := opendb.QueryContext(context.Background(), "SELECT VIDEOID,CHANNELID,ISO2,LAT,LONG,RADIUS,TAGID,DESCRIPTION,CITYID,COUNTRY,THUMBNAIL,PUBLISHEDAT FROM "+db.videoDBname+" WHERE ISO2=$1 LIMIT $2 ", ISO2, numberofvid)
	if err != nil {
		return nil, ReturnError(err)
	}
	return rows, nil
}

//Select one Random Video for Every country based on ISO2
func (db *Database) QueryRandomVideoFromEveryCountry() (*sql.Rows, error) {
	opendb, err := sql.Open("pgx", db.dsn.String())
	if err != nil {
		return nil, ReturnError(err)
	}
	defer func() {
		_ = opendb.Close()
		fmt.Println("closed")
	}()
	if err := opendb.PingContext(context.Background()); err != nil {
		return nil, ReturnError(err)

	}
	rows, err := opendb.QueryContext(context.Background(), "SELECT DISTINCT ON(ISO2) VIDEOID,CHANNELID,ISO2,LAT,LONG,RADIUS,TAGID,DESCRIPTION,CITYID,COUNTRY,THUMBNAIL,PUBLISHEDAT FROM "+db.videoDBname+" ORDER BY ISO2, RANDOM();")
	if err != nil {
		return nil, ReturnError(err)
	}
	return rows, nil
}
