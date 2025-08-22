package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/GenM4/go-ify/internal/responses"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	cfg := ApiConfig{
		ClientID:  os.Getenv("CLIENT_ID"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	cfg.Client = &http.Client{}
	cfg.getAccessKey()

	fmt.Println("Enter a spotify share URL:")
	var rawURL string
	_, err = fmt.Scanln(&rawURL)
	if err != nil {
		log.Fatal("Invalid input")
	}
	fmt.Println()

	asset := cfg.parseInput(rawURL)

	r := cfg.GetSpotifyAsset(asset)

	r.Log()

}

type SpotifyAsset struct {
	Type string
	ID   string
}

type ApiConfig struct {
	ClientID            string
	SecretKey           string
	Client              *http.Client
	AccessToken         string
	AccessTokenType     string
	AccessTokenDuration int
}

func (cfg *ApiConfig) parseInput(rawURL string) SpotifyAsset {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err)
	}

	cutPath := strings.Split(strings.TrimLeft(parsedURL.Path, "/"), "/")
	asset := SpotifyAsset{
		Type: cutPath[0],
		ID:   cutPath[1],
	}

	return asset
}

func (cfg *ApiConfig) GetSpotifyAsset(asset SpotifyAsset) responses.SpotifyType {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/" + asset.Type + "s/" + asset.ID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building " + asset.Type + " request")
	}
	req.Header.Add("Authorization", cfg.AccessTokenType+"  "+cfg.AccessToken)

	res, err := cfg.Client.Do(req)
	if err != nil {
		log.Print("Error sending " + asset.Type + " request")
	}

	var r responses.SpotifyType
	switch asset.Type {
	case "track":
		r = &responses.Track{}
	case "artist":
		r = &responses.Artist{}
	case "album":
		r = &responses.Album{}
	case "playlist":
		r = &responses.Playlist{}
	default:
		log.Fatalf("Unknown Spotify asset request: %s", asset.Type)
	}

	d := json.NewDecoder(res.Body)
	err = d.Decode(&r)
	if err != nil {
		log.Print("Error decoding " + asset.Type + " response")
	}

	return r

}

func (cfg *ApiConfig) getTrack(trackID string) responses.Track {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/tracks/" + trackID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building track request")
	}
	req.Header.Add("Authorization", cfg.AccessTokenType+"  "+cfg.AccessToken)

	res, err := cfg.Client.Do(req)
	if err != nil {
		log.Print("Error sending track request")
	}

	trackRes := responses.Track{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&trackRes)
	if err != nil {
		log.Print("Error decoding track response")
	}

	return trackRes
}

func (cfg *ApiConfig) getArtist(artistID string) responses.Artist {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/artists/" + artistID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building artist request")
	}
	req.Header.Add("Authorization", cfg.AccessTokenType+"  "+cfg.AccessToken)

	res, err := cfg.Client.Do(req)
	if err != nil {
		log.Print("Error sending artist request")
	}

	artistRes := responses.Artist{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&artistRes)
	if err != nil {
		log.Print("Error decoding artist token response")
	}

	return artistRes
}

func (cfg *ApiConfig) getAlbum(albumID string) responses.Album {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/albums/" + albumID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building album request")
	}
	req.Header.Add("Authorization", cfg.AccessTokenType+"  "+cfg.AccessToken)

	res, err := cfg.Client.Do(req)
	if err != nil {
		log.Print("Error sending album request")
	}

	albumRes := responses.Album{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&albumRes)
	if err != nil {
		log.Print("Error decoding album response")
	}

	return albumRes
}

func (cfg *ApiConfig) getPlaylist(playlistID string) responses.Playlist {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/playlists/" + playlistID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building playlist request")
	}
	req.Header.Add("Authorization", cfg.AccessTokenType+"  "+cfg.AccessToken)

	res, err := cfg.Client.Do(req)
	if err != nil {
		log.Print("Error sending playlist request")
	}

	playlistRes := responses.Playlist{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&playlistRes)
	if err != nil {
		log.Print("Error decoding playlist response")
	}

	return playlistRes
}

func (cfg *ApiConfig) getAccessKey() {
	tokenURL := "https://accounts.spotify.com/api/token"
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", cfg.ClientID)
	body.Set("client_secret", cfg.SecretKey)

	res, err := cfg.Client.PostForm(tokenURL, body)
	if err != nil {
		log.Print("Error with access token request")
	}

	type response struct {
		Token     string `json:"access_token"`
		TokenType string `json:"token_type"`
		ExpiresIn int    `json:"expires_in"`
	}

	accessKeyRes := response{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&accessKeyRes)
	if err != nil {
		log.Print("Error decoding access token response")
	}

	cfg.AccessToken = accessKeyRes.Token
	cfg.AccessTokenType = accessKeyRes.TokenType
	cfg.AccessTokenDuration = accessKeyRes.ExpiresIn

}
