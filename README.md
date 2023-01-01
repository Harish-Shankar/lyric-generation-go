
# Lyric Generation Go

I wanted to mess around with GoLang for a minute and also put a foot in the world of Natural Language Processing.

So this project combines both of those.

I have hopes to make this a fullstack application and a few more features
  - User can chose which albums to use
  - User can chose which songs to use
  - User can decide how may *verses* to generate

## Changing The Artist

To change the artist whose lyrics are being generated

  - Open ` getLyrics.go`
  - Change ` artist` on **Line 100**

## Deployment

To deploy this project run

```bash
  go run getLyrics.go
  go run markovChain.go
```

## Room For Improvement

Due to the way spotify works, it is possible that albums are repeated, leading to skewed results when generatinh lyrics
Additionally, sometimes the lyrics for the songs may not be accessible using the API, leading to similar issues.