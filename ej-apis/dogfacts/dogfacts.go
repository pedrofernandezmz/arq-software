package dogfacts

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Dogfacts []Dogfact
type Dogfact struct {
	Fact   string `json:"fact"`
}

func GetDogfacts(siteID string) (Dogfacts, error) {
	response, err := http.Get("https://dog-facts-api.herokuapp.com/api/v1/resources/dogs/all")
	dogfacts := Dogfacts{}
	if err != nil {
		err := errors.New("404")
		return dogfacts, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var facts Dogfacts
	err = json.Unmarshal(bytes, &facts)
	return facts, nil

}
