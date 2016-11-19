FROM nvidia/cuda:8.0-devel-ubuntu16.04

MAINTAINER Abdul Dakkak "dakkak@illinois.edu"

# Set one or more individual labels
LABEL com.webgpu.rai.version="0.0.1"
LABEL vendor="RAI Root Image"

#Build Essentials
RUN apt-get update && apt-get install --no-install-recommends -y  \
      software-properties-common                                  \
      python-software-properties &&                               \
    add-apt-repository ppa:ubuntu-toolchain-r/test &&             \
    apt-get update && apt-get install --no-install-recommends -y  \
      autoconf                                                    \
      automake                                                    \
      build-essential                                             \
      bzip2                                                       \
      ca-certificates                                             \
      clang                                                       \
      cmake                                                       \
      curl                                                        \
      file                                                        \
      gcc                                                         \
      git                                                         \
      git-core                                                    \
      libhdf5-10                                                  \
      libhdf5-10-dbg                                              \
      libhdf5-cpp-11                                              \
      libhdf5-cpp-11-dbg                                          \
      libhdf5-dev                                                 \
      libhdf5-openmpi-10                                          \
      libhdf5-openmpi-10-dbg                                      \
      libhdf5-openmpi-dev                                         \
      libopenmpi-dev                                              \
      make                                                        \
      mercurial                                                   \
      ntp                                                         \
      pkg-config                                                  \
      tar                                                         \
      unzip                                                       \
      wget                                                        \
      zlib1g-dbg &&                                               \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install --no-install-recommends -y  \
      linux-perf &&                                               \
    apt-get clean &&                                              \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* &&              \
    mv /usr/bin/perf_4.4 /usr/bin/perf

# RUN apt-get update && apt-get install --no-install-recommends -y  \
#       linux-cloud-tools-`uname -r`                                \
#       linux-tools-`uname -r`                                      \
#       linux-tools-common                                          \
#       linux-tools-generic &&                                      \
#     rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.7.3
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 508028aac0654e993564b6e2014bf2d4a9751e3b286661b0b0040046cf18028e

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN mkdir -p /build /src /data

WORKDIR /build
