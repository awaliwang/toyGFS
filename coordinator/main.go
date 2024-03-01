package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type file struct {
	Name   string `json:"filename"`
	Chunks int    `json:"chunks"`
}

var filemap = map[string][]int{
	"file1": {1, 2, 3},
	"file2": {4, 5},
}

var replicaLocations = map[int][]string{
	1: {"a", "b", "c"},
	2: {"a", "b", "c"},
	3: {"a", "b", "c"},
	4: {"a", "b", "c"},
	5: {"a", "b", "c"},
}

var chunkname = 6

// namespace, mapping from files to chunks,
// current locations of chunks
func main() {
	router := gin.Default()
	router.POST("/create", createFile)
	router.GET("/test", test)
	router.Run("localhost:8080")
}

func createFile(c *gin.Context) {
	// give each chunk a chunk handle
	var newFile file
	if err := c.BindJSON(&newFile); err != nil {
		return
	}
	var numberOfChunks = newFile.Chunks
	var chunks []int
	for i := 0; i < numberOfChunks; i++ {
		chunks = append(chunks, chunkname)
		chunkname += 1
	}
	// map file to chunks
	filemap[newFile.Name] = chunks
	// record chunk replica locations
	var writeLocations = map[int][]string{}
	for _, i := range chunks {
		replicaLocations[i] = []string{"a", "b", "c"}
		writeLocations[i] = []string{"a", "b", "c"}
	}
	// reply to client with locations of chunk servers to write replicas to
	c.IndentedJSON(http.StatusCreated, writeLocations)
}

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hello world")
}
