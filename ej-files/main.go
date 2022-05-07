package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	path = "./example.txt"
)

func main() {
	if err := CreateFile(path); err != nil {
		panic(err)
	}

	if err := WriteFile(path, []byte("Ejemplo de file --> Me llamo Pedro")); err != nil {
		panic(err)
	}

	bytes, err := ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Content of file: %s", string(bytes)))
}

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func WriteFile(path string, content []byte) error {
	if err := os.WriteFile(path, content, os.ModeAppend); err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
