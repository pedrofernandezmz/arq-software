//PKG.GO.DEV/NET/HTTP
package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	cats, err := GetCategories("MLA")
	if err != nil {
		//validar
	}
	fmt.Println("las categorias de MLA son...")
}

func GetCategories(siteID string) (Categories, error) {
	response := http.Get("https://api.mercadolibre.com/sites/MLA/categories")                    //completar
	bytes := ioutil.ReadAll(response.Bytes) //completar
	var cats Categories
	err := json.Unmarshal(bytes, &cats)
	return cats, nil

}

Dentro de la carpeta arq-spftware
git init
git remote add origin https:kjsfhalkjshfkhf
git status
git checkout -b wj-go-http
git branch?
ad commit etc
git push -u origin ej-go-http

