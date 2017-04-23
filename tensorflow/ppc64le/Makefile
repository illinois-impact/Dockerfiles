
.PHONY: all

all: build

push: build
	docker push webgpu/tensorflow:ppc64le
build:
	docker build . -t webgpu/tensorflow:ppc64le