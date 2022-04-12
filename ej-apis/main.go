package main

import (
	"fmt"
	"github.com/pedrofernandezmz/arq-software/ej-apis/dogfacts"
)

func main() {
	facts, err := dogfacts.GetDogfacts("dogs")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Los facts de la API sobre los perros es...")
	//for i := 0; i < len(facts); i++ {
	for i := 0; i < 5; i++ {
		fmt.Println("Fact:", facts[i].Fact, "\n")
	}
}