package service

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	CalledWithUrl string
	ReturnResp    *http.Response
	ReturnError   error
}

// Get implements HttpClient
func (c *MockClient) Get(url string) (resp *http.Response, err error) {
	c.CalledWithUrl = url
	return c.ReturnResp, c.ReturnError
}

var _ HttpClient = &MockClient{}
var _ Writer = &MockWriter{}

type MockWriter struct {
	writeMethodCalled bool
	fileName          string
	data              []byte
	perm              fs.FileMode
	ReturnError       error
}

// WriteFile implements Writer
func (w *MockWriter) WriteFile(name string, data []byte, perm fs.FileMode) error {
	w.writeMethodCalled = true
	w.fileName = name
	w.data = data
	w.perm = perm
	return w.ReturnError
}

func TestService(t *testing.T) {

	t.Run("should return error when there in error in making request", func(t *testing.T) {
		mockClient := MockClient{ReturnError: fmt.Errorf("error in backend")}

		err := ReadDataFromBackendAndWriteToFileV2("testUrl", "fileName", &mockClient, DefaultWriter{})

		assert.Equal(t, "testUrl", mockClient.CalledWithUrl)
		assert.EqualError(t, err, "error from backend: error in backend")
	})

	t.Run("should return response when there in no error", func(t *testing.T) {

		mockClient := MockClient{ReturnError: nil, ReturnResp: &http.Response{Body: io.NopCloser(strings.NewReader("Hello, world!"))}}
		mockWriter := &MockWriter{ReturnError: nil}
		err := ReadDataFromBackendAndWriteToFileV2("testUrl", "fileName", &mockClient, mockWriter)

		assert.Equal(t, "testUrl", mockClient.CalledWithUrl)
		assert.Nil(t, err)
		assert.True(t, mockWriter.writeMethodCalled)
		assert.Equal(t, "fileName", mockWriter.fileName)
		assert.Equal(t, fs.FileMode(0666), mockWriter.perm)
		assert.Equal(t, []byte("Hello, world!"), mockWriter.data)
	})

}
