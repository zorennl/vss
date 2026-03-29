package main

import (
	"fmt"
	"os"
)

func CreateFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(data)
	fmt.Println("Created file named ", filename)
}

func ReadFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func WriteFile(filename string, data string) {
	file, err := os.OpenFile(filename, 0644, os.FileMode(os.O_RDWR))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(data)
	fmt.Println("Write succesful")
}
