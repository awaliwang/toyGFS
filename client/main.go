package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type File struct {
	Filename string
	Chunks   int
}

func create(filename string, chunks int) File {
	newFile := File{filename, chunks}
	return newFile
}

func getWriteReplicas(writeFile File) []byte {
	url := "http://localhost:8080/create"
	jsonStr, err := json.Marshal(writeFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	return body
}

func main() {
	writeFile := create("file3", 3)

	chunkLocations := getWriteReplicas(writeFile)

	fmt.Println(chunkLocations)

}
