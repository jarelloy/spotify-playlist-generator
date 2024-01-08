package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/zmb3/spotify"
)

var genreList = []string{"acoustic", "afrobeat", "alt-rock", "alternative",
	"ambient", "anime", "black-metal", "bluegrass", "blues", "bossanova", "brazil",
	"breakbeat", "british", "cantopop", "chicago-house", "children", "chill",
	"classical", "club", "comedy", "country", "dance", "dancehall", "death-metal",
	"deep-house", "detroit-techno", "disco", "disney", "drum-and-bass", "dub",
	"dubstep", "edm", "electro", "electronic", "emo", "folk", "forro", "french",
	"funk", "garage", "german", "gospel", "goth", "grindcore", "groove", "grunge",
	"guitar", "happy", "hard-rock", "hardcore", "hardstyle", "heavy-metal",
	"hip-hop", "holidays", "honky-tonk", "house", "idm", "indian", "indie",
	"indie-pop", "industrial", "iranian", "j-dance", "j-idol", "j-pop", "j-rock",
	"jazz", "k-pop", "kids", "latin", "latino", "malay", "mandopop", "metal",
	"metal-misc", "metalcore", "minimal-techno", "movies", "mpb", "new-age",
	"new-release", "opera", "pagode", "party", "philippines-opm", "piano", "pop",
	"pop-film", "post-dubstep", "power-pop", "progressive-house", "psych-rock",
	"punk", "punk-rock", "r-n-b", "rainy-day", "reggae", "reggaeton", "road-trip",
	"rock", "rock-n-roll", "rockabilly", "romance", "sad", "salsa", "samba",
	"sertanejo", "show-tunes", "singer-songwriter", "ska", "sleep", "songwriter",
	"soul", "soundtracks", "spanish", "study", "summer", "swedish", "synth-pop",
	"tango", "techno", "trance", "trip-hop", "turkish", "work-out", "world-music"}

type SongInformation struct {
	Title       string        `json:"title"`
	Album       string        `json:"album"`
	Duration    time.Duration `json:"duration"`
	Artist      string        `json:"artist"`
	PlaybackURI string        `json:"playbackuri"`
	URI         spotify.ID    `json:"uri"`
}

type PlaylistTracks struct {
	SelectedSongs []spotify.ID `json:"selectedSongs"`
}

type SpotifyClient struct {
	Client spotify.Client
	UserID string
}

func SearchByGenre(genre string, client *spotify.Client) []SongInformation {
	var songInfoList []SongInformation
	tracks, err := client.Search("genre\":"+genre, spotify.SearchTypePlaylist)
	if err != nil {
		fmt.Println(err)
	}

	playlistId := tracks.Playlists.Playlists[0].ID
	playlistTracks, err := client.GetPlaylistTracks(playlistId)
	if err != nil {
		fmt.Println(err)
	}

	for _, track := range playlistTracks.Tracks {
		songInfo := SongInformation{
			Title:       track.Track.Name,
			Album:       track.Track.Album.Name,
			Duration:    track.Track.TimeDuration(),
			Artist:      GetArtistsNames(track.Track.Artists),
			PlaybackURI: string(track.Track.URI),
			URI:         track.Track.ID,
		}
		songInfoList = append(songInfoList, songInfo)
	}

	return songInfoList
}

func GetArtistsNames(artists []spotify.SimpleArtist) string {
	var names []string
	for _, artist := range artists {
		names = append(names, artist.Name)
	}
	return fmt.Sprintf(strings.Join(names, ", "))
}

func CreatePlaylist(genre string, sc *SpotifyClient, tracksURI PlaylistTracks) error {
	currentDate := time.Now().Format("2006-01-02")
	playlistName := fmt.Sprintf("%s : %s Genre Playlist", currentDate, genre)
	playlist, err := sc.Client.CreatePlaylistForUser(sc.UserID, playlistName, "", true)
	if err != nil {
		fmt.Println("Error creating playlist")
		return err
	}
	_, err = sc.Client.AddTracksToPlaylist(playlist.ID, tracksURI.SelectedSongs...)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func RandomGenreSelector() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(genreList))
	return genreList[randomIndex]
}
