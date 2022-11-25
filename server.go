package main

import (
	"encoding/json"
	"fmt"
	"guess_music_youtube/internal/controler"
	"guess_music_youtube/internal/controler/youtubevIdeoupdater/updater"
	Database "guess_music_youtube/internal/database"
	"net/http"
	"os"
	"strconv"
	"time"

	. "guess_music_youtube/internal/auxfunctions"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var db Database.Database
var control controler.Controler
var videoUpdater updater.VideoUpdater

func init() {
	/*err := db.CreateDBS()
	if err != nil {
		fmt.Println(err)
	}*/
}

func main() {
	time.Sleep(time.Second * 10)
	err := db.InitDb()
	if err != nil {
		ReturnError(err)
	}
	videoPerDay, err := strconv.ParseInt(os.Getenv("SEARCH_VID_PERDAY"), 10, 64)
	if err != nil {
		ReturnError(err)
	}
	UpdateDays, err := strconv.ParseInt(os.Getenv("SEARCH_VID_DAYS"), 10, 64)
	if err != nil {
		ReturnError(err)
	}

	videoUpdater.Init(int(videoPerDay), int(UpdateDays), db)
	go videoUpdater.UpdaterMain()
	control = control.NewControler(db)

	r := mux.NewRouter()
	r.HandleFunc("/game/{countryid}/{numberofvideos}", GetGameCountry).
		Methods("GET")
	r.HandleFunc("/game/everycountryvideo", GetRandomVideoForEveryCountry).
		Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", r)
}

/*Get Country By Iso2 and number*/
func GetGameCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Print("yfonfodfho")
	vars := mux.Vars(r)
	countryId := vars["countryid"]
	numberofvid, err := strconv.Atoi(vars["numberofvideos"])
	if err != nil {
		fmt.Println(err)
	}
	youtubevideos, err := control.GetCountryVideos(numberofvid, countryId)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(youtubevideos); err != nil {
		fmt.Println(err)
	}
	return
}

//Get Random Video From every Country ,that way the web server is able to create a cache for the client
func GetRandomVideoForEveryCountry(w http.ResponseWriter, r *http.Request) {
	youtubevideos, err := control.GetRandomEveryCountry()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(youtubevideos); err != nil {
		fmt.Println(err)
	}
	return
}
