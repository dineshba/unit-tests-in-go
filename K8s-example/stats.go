package stats

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	clientPkg "sigs.k8s.io/controller-runtime/pkg/client"
)

func GetConfigMapCount(ctx context.Context, client clientPkg.Client, ns string) (int, error) {
	cmList := corev1.ConfigMapList{}
	opts := []clientPkg.ListOption{
		clientPkg.InNamespace("default"),
	}

	err := client.List(ctx, &cmList, opts...)
	if err != nil {
		return 0, fmt.Errorf("error getting list of configmaps: %s", err.Error())
	}

	return len(cmList.Items), nil
}
