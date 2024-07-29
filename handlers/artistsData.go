package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Res struct {
	ID           int      `json:"id"`
	Img          string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var ArtistsObj []Res

func GetArtistsData() {
	const URL = "https://groupietrackers.herokuapp.com/api/artists"

	// query api
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	////read in data from the incoming byte stream
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// json.NewDecoder(res).Decode(any)

	json.Unmarshal(resData, &ArtistsObj)
}
