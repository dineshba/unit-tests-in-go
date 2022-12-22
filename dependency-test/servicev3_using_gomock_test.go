package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mocks "unittest/mocks"
)

func TestServiceV3(t *testing.T) {

	t.Run("should return error when there in error in making request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClient := mocks.NewMockHttpClient(ctrl)
		mockWriter := mocks.NewMockWriter(ctrl)
		mockClient.EXPECT().Get("testurl").Return(nil, fmt.Errorf("network error"))

		s := NewService(mockClient, mockWriter)
		err := s.ReadDataFromBackendAndWriteToFile("testurl", "testFile")

		assert.Equal(t, fmt.Errorf("error from backend: network error"), err)
	})

	t.Run("should return response when there in no error", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClient := mocks.NewMockHttpClient(ctrl)
		mockWriter := mocks.NewMockWriter(ctrl)
		mockClient.EXPECT().Get("testurl").Return(&http.Response{Body: io.NopCloser(strings.NewReader("Hello, world!"))}, nil)
		mockWriter.EXPECT().WriteFile("testFile", []byte("Hello, world!"), os.FileMode(0666)).Times(1)

		s := NewService(mockClient, mockWriter)
		err := s.ReadDataFromBackendAndWriteToFile("testurl", "testFile")

		assert.Nil(t, err)
	})

}
