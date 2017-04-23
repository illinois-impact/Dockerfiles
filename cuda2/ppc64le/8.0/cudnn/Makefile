
.PHONY: all

all: build

push: build
	docker push webgpu/cuda_cudnn:ppc64le_8.0
build:
	docker build . -t webgpu/cuda_cudnn:ppc64le_8.0