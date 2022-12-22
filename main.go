package main

import (
	"fmt"
	"time"
	service "unittest/dependency-test"
	_ "unittest/simple-test"
)

func main() {
	service.ReadDataFromBackendAndWriteToFile("https://dummyjson.com/products/1", fmt.Sprintf("output-%d.json", time.Now().Unix()))
}
