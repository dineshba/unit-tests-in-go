package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadDataFromBackendAndWriteToFile(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error from backend: %s", err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body from response: %s", err)
	}

	err = os.WriteFile(fileName, data, 0666)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err.Error())
	}

	return nil
}
