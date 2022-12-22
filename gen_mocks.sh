#!/bin/bash

set -e

rm -rf mocks
mkdir -p mocks

mockgen -source=dependency-test/servicev2.go -destination mocks/service.go