package services

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test1-tribal/config"
)

var Config, _ = config.LoadConfig("./")

type ResponseClientApple struct {
	ResultCount int              `json:"resultCount"`
	Results     []map[string]any `json:"results"`
}

type ResponseClientChartlyrics struct {
	XMLName           xml.Name `xml:"ArrayOfSearchLyricResult"`
	Text              string   `xml:",chardata"`
	Xsd               string   `xml:"xsd,attr"`
	Xsi               string   `xml:"xsi,attr"`
	Xmlns             string   `xml:"xmlns,attr"`
	SearchLyricResult []struct {
		Text          string `xml:",chardata"`
		Nil           string `xml:"nil,attr"`
		TrackId       string `xml:"TrackId"`
		LyricChecksum string `xml:"LyricChecksum"`
		LyricId       string `xml:"LyricId"`
		SongUrl       string `xml:"SongUrl"`
		ArtistUrl     string `xml:"ArtistUrl"`
		Artist        string `xml:"Artist"`
		Song          string `xml:"Song"`
		SongRank      string `xml:"SongRank"`
		TrackChecksum string `xml:"TrackChecksum"`
	} `xml:"SearchLyricResult"`
}

func GetSongOfAllClients(name string, artist string, album string) (merge []any) {

	/*
		🔴😥
		No se puede realizar busquedas de diferentes entidades
		ejemplo correcto term=jack+johnson&entity=musicVideo
	*/
	var filterApple string = "term=" + name

	/*
		🔴😥
		Se necesitan tanto el nombre del artista como el título de la canción para filterChartlyrics.
		Las palabras vacías se eliminan de la consulta de búsqueda

		y algunos datos requeridos no estan en el xml
	*/
	var filterChartlyrics string = "song=" + name + "&artist=" + artist

	if len(name) > 0 {
		filterApple = "term=" + name
	}

	if len(artist) > 0 {
		filterApple = filterApple + ""
	}

	if len(album) > 0 {
		filterApple = filterApple + ""
	}

	println("filter Apple : ", filterApple)
	println("filter Chartlyrics : ", filterChartlyrics)
	dataChartlyrics := GetSongClientChartlyrics(filterChartlyrics)
	dataApple := GetSongClientApple(filterApple)

	merge = append(dataChartlyrics, dataApple...)
	fmt.Println(len(merge))

	return

}

func GetSongClientApple(url string) (responseSongs []any) {

	var responseObject = getSongClientApi(Config.ClientAppleApi + url)

	for i := 0; i < responseObject.ResultCount; i++ {
		mapa := make(map[string]interface{})
		song := responseObject.Results[i]
		mapa["IdSong"] = song["trackId"]
		mapa["Name"] = song["trackName"]
		mapa["Artist"] = song["artistName"]
		mapa["Duration"] = song["trackTimeMillis"]
		mapa["Album"] = song["collectionName"]
		mapa["Artwork"] = song["previewUrl"]
		mapa["Price"] = song["trackPrice"]
		mapa["Origin"] = "Apple"

		responseSongs = append(responseSongs, mapa)
		/* responseSongs = append(responseSongs, models.SongResponse{
			IdSong:   responseObject.Results[i]["trackId"],
			Name:     responseObject.Results[i]["trackName"],
			Duration: responseObject.Results[i]["trackTimeMillis"],
			Album:    responseObject.Results[i]["trackCensoredName"],
			Artwork:  responseObject.Results[i]["primaryGenreName"],
			Price:    responseObject.Results[i]["trackPrice"],
			Origin:   "Apple",
		})
		fmt.Println(responseObject.Results[i]["artistName"])*/
	}
	return
}

func GetSongClientChartlyrics(url string) (responseSongs []any) {

	var responseObject = getSongClientSoap(Config.ClientChartlyricsApi + url)

	for i := 0; i < len(responseObject.SearchLyricResult)-1; i++ {
		mapa := make(map[string]interface{})
		song := responseObject.SearchLyricResult[i]
		mapa["IdSong"] = song.TrackId
		mapa["Name"] = song.Song
		mapa["Artist"] = song.Artist
		mapa["Duration"] = song.SongRank
		mapa["Album"] = song.Song
		mapa["Artwork"] = song.SongUrl
		mapa["Price"] = song.SongRank
		mapa["Origin"] = "chartlyrics"

		responseSongs = append(responseSongs, mapa)
	}
	return
}

func getSongClientApi(url string) (responseObject ResponseClientApple) {

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(responseData, &responseObject)

	return
}

func getSongClientSoap(url string) (responseObject ResponseClientChartlyrics) {

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	html, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	website, err := ioutil.ReadAll(html.Body)
	if err != nil {
		log.Fatal(err)
	}
	html.Body.Close()

	xml.Unmarshal(website, &responseObject)

	return
}
