package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Repository interface {
	GetTrack(id string) (*Track, error)
	GetArtist(id string) (*Artist, error)
	GetAlbum(id string) (*Album, error)
	GetPlaylist(id string) (*Playlist, error)
}

type SpotifyRepository struct {
	Repository
	AccessToken         string
	AccessTokenDuration int
}

func NewSpotifyRepository() (*SpotifyRepository, error) {
	token, dur, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	return &SpotifyRepository{
		AccessToken:         token,
		AccessTokenDuration: dur,
	}, nil
}

func getAccessToken() (token string, duration int, err error) {
	cid, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		return "", 0, errors.New("Environment variable 'CLIENT_ID' does not exist")
	}

	secret, exists := os.LookupEnv("SECRET_KEY")
	if !exists {
		return "", 0, errors.New("Environment variable 'SECRET_KEY' does not exist")
	}

	tokenURL := "https://accounts.spotify.com/api/token"
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", cid)
	body.Set("client_secret", secret)

	client := http.Client{}
	res, err := client.PostForm(tokenURL, body)
	if err != nil {
		log.Print("Error with access token request")
		return "", 0, err
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
		return "", 0, err
	}

	token = accessKeyRes.Token
	duration = accessKeyRes.ExpiresIn

	return

}

func (repo *SpotifyRepository) GetTrack(id string) (*Track, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/tracks/" + id,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building track request")
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer  "+repo.AccessToken)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print("Error sending track request")
		return nil, err
	}

	track := Track{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&track)
	if err != nil {
		log.Print("Error decoding track response")
		return nil, err
	}

	return &track, nil
}

func (repo *SpotifyRepository) GetArtist(id string) (*Artist, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/artists/" + id,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building artist request")
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer  "+repo.AccessToken)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print("Error sending artist request")
		return nil, err
	}

	artist := Artist{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&artist)
	if err != nil {
		log.Print("Error decoding artist token response")
		return nil, err
	}

	return &artist, nil
}

func (repo *SpotifyRepository) GetAlbum(id string) (*Album, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/albums/" + id,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building album request")
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer  "+repo.AccessToken)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print("Error sending album request")
		return nil, err
	}

	album := Album{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&album)
	if err != nil {
		log.Print("Error decoding album response")
		return nil, err
	}

	return &album, nil
}

func (repo *SpotifyRepository) GetPlaylist(id string) (*Playlist, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "api.spotify.com",
		Path:   "v1/playlists/" + id,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		log.Print("Error building playlist request")
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer  "+repo.AccessToken)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print("Error sending playlist request")
		return nil, err
	}

	playlist := Playlist{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&playlist)
	if err != nil {
		log.Print("Error decoding playlist response")
		return nil, err
	}

	return &playlist, nil
}
