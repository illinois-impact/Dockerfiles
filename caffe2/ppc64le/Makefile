
.PHONY: all

all: build

push: build
	docker push webgpu/caffe2:ppc64le
build:
	docker build . -t webgpu/caffe2:ppc64le