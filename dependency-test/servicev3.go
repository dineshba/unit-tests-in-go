package service

import (
	"fmt"
	"io/ioutil"
)

type Service struct {
	client HttpClient
	writer Writer
}

func NewService(client HttpClient, writer Writer) Service {
	return Service{client: client, writer: writer}
}

func (s Service) ReadDataFromBackendAndWriteToFile(url, fileName string) error {
	resp, err := s.client.Get(url)
	if err != nil {
		return fmt.Errorf("error from backend: %s", err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body from response: %s", err)
	}

	err = s.writer.WriteFile(fileName, data, 0666)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err.Error())
	}

	return nil
}
