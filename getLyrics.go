package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func getLyrics() {
	var err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	var CLIENT_ID = os.Getenv("CLIENT_ID")
	var CLIENT_SECRET = os.Getenv("CLIENT_SECRET")

	var ctx = context.Background()
	var config = &clientcredentials.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
}
