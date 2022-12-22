package stats

import (
	"context"
	"testing"

	mock "unittest/mocks/client"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigMapCount(t *testing.T) {
	t.Run("when there is error from client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.TODO()
		ns := "default"
		mockClient := mock.NewMockClient(ctrl)
		// mockClientSetup

		count, err := GetConfigMapCount(ctx, mockClient, ns)

		assert.Equal(t, 0, count)
		assert.EqualError(t, err, "error getting list of configmaps: ns not found")
	})

	t.Run("when there is no error from client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.TODO()
		ns := "default"
		mockClient := mock.NewMockClient(ctrl)
		// mockClientSetup

		count, err := GetConfigMapCount(ctx, mockClient, ns)

		assert.Equal(t, 3, count)
		assert.Nil(t, err)
	})
}
