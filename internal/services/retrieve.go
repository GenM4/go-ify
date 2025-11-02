package services

import (
	"net/url"
	"strings"

	"github.com/GenM4/go-ify/internal/api"
)

type Retriever interface {
	GetTrack(url string) (*api.Track, error)
	GetArtist(url string) (*api.Artist, error)
	GetAlbum(url string) (*api.Album, error)
	GetPlaylist(url string) (*api.Playlist, error)
}

type SpotifyRetriever struct {
	Retriever
	Repo *api.SpotifyRepository
}

func NewSpotifyRetriever(repo *api.SpotifyRepository) *SpotifyRetriever {
	return &SpotifyRetriever{Repo: repo}
}

func (ret *SpotifyRetriever) GetTrack(url string) (*api.Track, error) {
	_, id, err := parseInput(url)
	if err != nil {
		return nil, err
	}

	track, err := ret.Repo.GetTrack(id)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (ret *SpotifyRetriever) GetArtist(url string) (*api.Artist, error) {
	_, id, err := parseInput(url)
	if err != nil {
		return nil, err
	}

	artist, err := ret.Repo.GetArtist(id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
func (ret *SpotifyRetriever) GetAlbum(url string) (*api.Album, error) {
	_, id, err := parseInput(url)
	if err != nil {
		return nil, err
	}

	album, err := ret.Repo.GetAlbum(id)
	if err != nil {
		return nil, err
	}

	return album, nil
}
func (ret *SpotifyRetriever) GetPlaylist(url string) (*api.Playlist, error) {
	_, id, err := parseInput(url)
	if err != nil {
		return nil, err
	}

	playlist, err := ret.Repo.GetPlaylist(id)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func parseInput(rawURL string) (assetType, id string, err error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}

	cutPath := strings.Split(strings.TrimLeft(parsedURL.Path, "/"), "/")
	assetType = cutPath[0]
	id = cutPath[1]

	return
}

/*
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
*/
