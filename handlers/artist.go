package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type locations struct {
	LocationsD []string `json:"locations"`
}
type dates struct {
	DatesD []string `json:"dates"`
}
type relation struct {
	RelationD map[string][]string `json:"datesLocations"`
}

type Data struct {
	Old      Res
	Loc      locations
	Date     dates
	Relation relation
}

var (
	newlocation locations
	Dates       dates
	Relation    relation
)

func GetArtist(res http.ResponseWriter, req *http.Request) {
	file := "artist.html"

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println(err)
	}

	elem := strings.Split(string(req.URL.Path), "/")
	id, _ := strconv.Atoi(elem[len(elem)-1])

	getDates(id, "https://groupietrackers.herokuapp.com/api/locations/", "locations")
	getDates(id, "https://groupietrackers.herokuapp.com/api/dates/", "dates")
	getDates(id, "https://groupietrackers.herokuapp.com/api/relation/", "relation")
	if id > 51 || id < 0 {
		http.NotFound(res, req)
		return
	}
	result := Data{
		Old:      ArtistsObj[id-1],
		Loc:      newlocation,
		Date:     Dates,
		Relation: Relation,
	}

	err = tmpl.Execute(res, result)
	if err != nil {
		fmt.Println(err)
	}
}

func getDates(id int, url string, info string) {
	DatesURL := fmt.Sprint(url, id)

	DatesData, err := http.Get(DatesURL)
	if err != nil {
		fmt.Println(err)
	}

	content, err1 := io.ReadAll(DatesData.Body)

	if err1 != nil {
		fmt.Println(err1)
	}
	if info == "locations" {
		json.Unmarshal(content, &newlocation)
	} else if info == "dates" {
		json.Unmarshal(content, &Dates)
	} else if info == "relation" {
		json.Unmarshal(content, &Relation)
	}
}
