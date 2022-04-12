//PKG.GO.DEV/NET/HTTP
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Dogfacts []Dogfact
type Dogfact struct {
	//type Categories []struct {
	Fact   string `json:"fact"`
	//Name string `json:"name"`
}

func main() {
	facts, err := GetDogfacts("dogs")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Las categorias de MLA son...")
	for i := 0; i < len(facts); i++ {
		fmt.Println("Fact:", facts[i].Fact, "\n")
		//fmt.Println("Name:", cats[i].Name, "\n")
	}
}

func GetDogfacts(siteID string) (Dogfacts, error) {
	response, err := http.Get("https://dog-facts-api.herokuapp.com/api/v1/resources/dogs/all")
	dogfacts := Dogfacts{}
	if err != nil {
		err := errors.New("404")
		return dogfacts, err
	} //completar
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	} //completar
	var facts Dogfacts
	err = json.Unmarshal(bytes, &facts)
	return facts, nil

}
