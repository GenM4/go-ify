package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/GenM4/go-ify/internal/responses"
)

type SpotifyAsset struct {
	Type string
	ID   string
}

type SpotifyApi struct {
	Client              *http.Client
	AccessToken         string
	AccessTokenType     string
	AccessTokenDuration int
}

func SpotifyApiInit() *SpotifyApi {
	api := &SpotifyApi{
		Client: &http.Client{},
	}

	if err := api.GetAccessKey(); err != nil {
		log.Fatal(err)
	}

	return api
}

func (api *SpotifyApi) parseInput(rawURL string) SpotifyAsset {
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

func (api *SpotifyApi) GetSpotifyAsset(rawURL string) responses.SpotifyType {
	asset := api.parseInput(rawURL)

	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/" + asset.Type + "s/" + asset.ID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building " + asset.Type + " request")
	}
	req.Header.Add("Authorization", api.AccessTokenType+"  "+api.AccessToken)

	res, err := api.Client.Do(req)
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

func (api *SpotifyApi) GetTrack(rawURL string) responses.Track {
	asset := api.parseInput(rawURL)

	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/" + asset.Type + "s/" + asset.ID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building track request")
	}
	req.Header.Add("Authorization", api.AccessTokenType+"  "+api.AccessToken)

	res, err := api.Client.Do(req)
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

func (api *SpotifyApi) GetArtist(rawURL string) responses.Artist {
	asset := api.parseInput(rawURL)

	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/" + asset.Type + "s/" + asset.ID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building artist request")
	}
	req.Header.Add("Authorization", api.AccessTokenType+"  "+api.AccessToken)

	res, err := api.Client.Do(req)
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

func (api *SpotifyApi) GetAlbum(rawURL string) responses.Album {
	asset := api.parseInput(rawURL)

	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/" + asset.Type + "s/" + asset.ID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building album request")
	}
	req.Header.Add("Authorization", api.AccessTokenType+"  "+api.AccessToken)

	res, err := api.Client.Do(req)
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

func (api *SpotifyApi) GetPlaylist(rawURL string) responses.Playlist {
	asset := api.parseInput(rawURL)

	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/" + asset.Type + "s/" + asset.ID,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building playlist request")
	}
	req.Header.Add("Authorization", api.AccessTokenType+"  "+api.AccessToken)

	res, err := api.Client.Do(req)
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

func (api *SpotifyApi) GetAccessKey() error {
	cid, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		return errors.New("Environment variable 'CLIENT_ID' does not exist")
	}

	secret, exists := os.LookupEnv("SECRET_KEY")
	if !exists {
		return errors.New("Environment variable 'SECRET_KEY' does not exist")
	}

	tokenURL := "https://accounts.spotify.com/api/token"
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", cid)
	body.Set("client_secret", secret)

	res, err := api.Client.PostForm(tokenURL, body)
	if err != nil {
		log.Print("Error with access token request")
		return err
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
		return err
	}

	api.AccessToken = accessKeyRes.Token
	api.AccessTokenType = accessKeyRes.TokenType
	api.AccessTokenDuration = accessKeyRes.ExpiresIn

	return nil

}
