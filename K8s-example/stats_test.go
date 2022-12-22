package stats

import (
	"context"
	"fmt"
	"testing"

	mock "unittest/mocks/client"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	clientPkg "sigs.k8s.io/controller-runtime/pkg/client"
)

func TestGetConfigMapCount(t *testing.T) {
	t.Run("when there is error from client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.TODO()
		ns := "default"
		mockClient := mock.NewMockClient(ctrl)
		// mockClientSetup
		mockClient.EXPECT().List(ctx, gomock.Any(), []clientPkg.ListOption{
			clientPkg.InNamespace(ns),
		}).Return(fmt.Errorf("ns not found"))

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
		mockClient.EXPECT().List(ctx, gomock.Any(), []clientPkg.ListOption{
			clientPkg.InNamespace(ns),
		}).DoAndReturn(func(_, listInterface, _ interface{}) error {
			list := listInterface.(*corev1.ConfigMapList)
			list.Items = []corev1.ConfigMap{
				{},
				{},
				{},
			}
			return nil
		})

		count, err := GetConfigMapCount(ctx, mockClient, ns)

		assert.Equal(t, 3, count)
		assert.Nil(t, err)
	})
}
