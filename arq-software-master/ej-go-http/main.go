package main

import (
	"fmt"
	"github.com/emikohmann/arq-software/ej-go-http/categories"
	"os"
)

func main() {
	categories, err := categories.GetCategories("MLA")
	if err != nil {
		fmt.Println("Error getting categories: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Categorias obtenidas de la API de Mercado Libre:")
	for _, category := range categories {
		fmt.Println(category.String())
	}
}
