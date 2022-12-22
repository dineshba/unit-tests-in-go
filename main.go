package main

import (
	"fmt"
	"net/http"
	"time"
	service "unittest/dependency-test"
	_ "unittest/simple-test"
)

func main() {
	service.ReadDataFromBackendAndWriteToFile("https://dummyjson.com/products/1", fmt.Sprintf("output-%d.json", time.Now().Unix()))
	service.ReadDataFromBackendAndWriteToFileV2(
		"https://dummyjson.com/products/1",
		fmt.Sprintf("outputv2-%d.json", time.Now().Unix()),
		&http.Client{},
		service.DefaultWriter{},
	)
}
