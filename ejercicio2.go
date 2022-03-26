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

type Categories []Category
type Category struct {
	//type Categories []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	cats, err := GetCategories("MLA")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Las categorias de MLA son...")
	for i := 0; i < len(cats); i++ {
		fmt.Println("Id:", cats[i].ID)
		fmt.Println("Name:", cats[i].Name, "\n")
	}
}

func GetCategories(siteID string) (Categories, error) {
	response, err := http.Get("https://api.mercadolibre.com/sites/MLA/categories")
	categories := Categories{}
	if err != nil {
		err := errors.New("404")
		return categories, err
	} //completar
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	} //completar
	var cats Categories
	err = json.Unmarshal(bytes, &cats)
	return cats, nil

}
