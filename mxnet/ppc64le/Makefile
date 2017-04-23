
.PHONY: all

all: build

push: build
	docker push webgpu/mxnet:ppc64le
build:
	docker build . -t webgpu/mxnet:ppc64le