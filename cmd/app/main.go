package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

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

	cfg.getAccessKey()

	client := http.Client{}

	fmt.Println("Enter a spotify artist URL:")
	var artistURL string
	_, err = fmt.Scanln(&artistURL)

	req, err := http.NewRequest(http.MethodGet, artistURL, http.NoBody)
	if err != nil {
		log.Print("Error building artist request")
	}
	req.Header.Add("Authorization", cfg.AccessTokenType+"  "+cfg.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		log.Print("Error sending artist request")
	}

	type response struct {
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
	}

	artistRes := response{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&artistRes)
	if err != nil {
		log.Print("Error decoding artist token response")
	}

	fmt.Println(artistRes.Name)
	fmt.Println(artistRes.Popularity)
	fmt.Println(artistRes.Type)

}

type ApiConfig struct {
	ClientID            string
	SecretKey           string
	AccessToken         string
	AccessTokenType     string
	AccessTokenDuration int
}

func (cfg *ApiConfig) getAccessKey() {
	tokenURL := "https://accounts.spotify.com/api/token"
	body := url.Values{}
	body.Set("grant_type", "client_credentials")
	body.Set("client_id", cfg.ClientID)
	body.Set("client_secret", cfg.SecretKey)

	res, err := http.PostForm(tokenURL, body)
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
