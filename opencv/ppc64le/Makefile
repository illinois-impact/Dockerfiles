
.PHONY: all

all: build

push: build
	docker push webgpu/opencv:ppc64le_3.2.0
build:
	docker build . -t webgpu/opencv:ppc64le_3.2.0