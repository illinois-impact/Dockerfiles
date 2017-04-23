
.PHONY: all

all: build

push: build
	docker push webgpu/cuda_runtime:ppc64le_8.0
build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" cross_build.go
	docker build . -t webgpu/cuda_runtime:ppc64le_8.0