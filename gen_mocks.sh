#!/bin/bash

set -e

rm -rf mocks
mkdir -p mocks

mockgen -source=dependency-test/servicev2.go -destination mocks/service.go

rm -rf mocks/client/
mkdir -p mocks/client/
mockgen sigs.k8s.io/controller-runtime/pkg/client Client >> mocks/client/client.go