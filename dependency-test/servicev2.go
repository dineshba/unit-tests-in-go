package service

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type Writer interface {
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type DefaultWriter struct{}

// WriteFile implements Writer
func (DefaultWriter) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

var _ Writer = DefaultWriter{}

func ReadDataFromBackendAndWriteToFileV2(url, fileName string, client HttpClient, writer Writer) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error from backend: %s", err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body from response: %s", err)
	}

	err = writer.WriteFile(fileName, data, 0666)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err.Error())
	}

	return nil
}
