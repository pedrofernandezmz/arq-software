package categories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (category Category) String() string {
	return fmt.Sprintf(" - %s: %s", category.ID, category.Name)
}

func GetCategories(siteID string) (Categories, error) {
	response, err := http.Get(getURL(siteID))
	if err != nil {
		return Categories{}, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Categories{}, err
	}
	return parseCategories(bytes)
}

func getURL(siteID string) string {
	return fmt.Sprintf("https://api.mercadolibre.com/sites/%s/categories", siteID)
}

func parseCategories(bytes []byte) (Categories, error) {
	var cats Categories
	if err := json.Unmarshal(bytes, &cats); err != nil {
		return Categories{}, err
	}
	return cats, nil
}
