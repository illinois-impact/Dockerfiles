FROM webgpu/cuda:ppc64le-8.0-devel
MAINTAINER Carl Pearson <pearson@illinois.edu>

RUN [ "cross-build-start" ]

RUN apt-get update && apt-get install -y --no-install-recommends \
        libmpich-dev mpich g++-5 && \
    rm -rf /var/lib/apt/lists/*

RUN [ "cross-build-end" ]

