package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func getAllAlbums(ctx context.Context, artist string, client spotify.Client) []spotify.ID {
	results, err := client.Search(ctx, artist, spotify.SearchTypeArtist)
	if err != nil {
		panic(err)
	}
	var artist_id spotify.ID
	if results.Artists != nil {
		for _, item := range results.Artists.Artists {
			if item.Name == artist {
				artist_id = item.ID
				break
			}
		}
	}

	artist_albums, err := client.GetArtistAlbums(ctx, artist_id, []spotify.AlbumType{1}, spotify.Market("US") /* , spotify.Limit(50) */)
	if err != nil {
		panic(err)
	}
	var album_ids []spotify.ID
	var album_names []string
	for _, album := range artist_albums.Albums {
		album_ids = append(album_ids, album.ID)
		album_names = append(album_names, album.Name)
	}

	for i := range album_names {
		fmt.Println(album_names[i])
	}

	return album_ids
}

func getAllSongs(ctx context.Context, albumIDList []spotify.ID, client spotify.Client) []string {
	var songs []string
	var songids []spotify.ID

	for _, album_id := range albumIDList {
		album_tracks, err := client.GetAlbumTracks(ctx, album_id)
		if err != nil {
			panic(err)
		}
		for _, track := range album_tracks.Tracks {
			if !contains(songids, track.ID) {
				songs = append(songs, track.Name)
				songids = append(songids, track.ID)
			}
		}
	}
	return songs
}

func contains(s []spotify.ID, e spotify.ID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

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

	artist := "Kendrick Lamar"
	albumIDList := getAllAlbums(ctx, artist, *client)
	songList := getAllSongs(ctx, albumIDList, *client)
}
