package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"spotify-go/utils"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	"github.com/zmb3/spotify"
)

var (
	clientID     string
	clientSecret = "your_client_secret"
	redirectURI  = "http://localhost:8080/callback"
	state        = "state"

	spotifyAuthenticator = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail, spotify.ScopePlaylistModifyPublic)
	Ch                   = make(chan *spotify.Client)
	client               spotify.Client
	sc                   utils.SpotifyClient
	genre                string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		fmt.Println("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		os.Exit(1)
	}
	spotifyAuthenticator.SetAuthInfo(clientID, clientSecret)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/callback", CallbackHandler)
	r.HandleFunc("/home", RandomizeGenrePageHandler)
	r.HandleFunc("/display-songs", DisplaySongsFromPlaylistHandler)
	r.HandleFunc("/save-songs", SaveSongsToPlaylistHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "templates/login.html", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := spotifyAuthenticator.AuthURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func RandomizeGenrePageHandler(w http.ResponseWriter, r *http.Request) {
	randomGenre := utils.RandomGenreSelector()
	RenderTemplate(w, "templates/randomizing.html", randomGenre)
}

func DisplaySongsFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	genre = r.URL.Query().Get("genre")
	songs := utils.SearchByGenre(genre, &client)
	RenderTemplate(w, "templates/song-list.html", songs)
}

func SaveSongsToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload utils.PlaylistTracks
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	utils.CreatePlaylist(genre, &sc, payload)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Selection received successfully"))
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyAuthenticator.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	client = spotifyAuthenticator.NewClient(token)
	Ch <- &client
	user, _ := client.CurrentUser()
	userId := user.ID

	sc = utils.SpotifyClient{
		Client: client,
		UserID: userId,
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
