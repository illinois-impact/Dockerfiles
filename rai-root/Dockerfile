FROM nvidia/cuda:8.0-cudnn5-devel-ubuntu16.04

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
      strace                                                      \
      tar                                                         \
      time                                                        \
      unzip                                                       \
      wget                                                        \
      zlib1g-dbg &&                                               \
    rm -rf /var/lib/apt/lists/*

# RUN apt-get update && apt-get install --no-install-recommends -y  \
#       linux-perf &&                                               \
#     apt-get clean &&                                              \
#     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* &&              \
#     mv /usr/bin/perf_4.4 /usr/bin/perf

# RUN apt-get update && apt-get install --no-install-recommends -y  \
#       linux-cloud-tools-`uname -r`                                \
#       linux-tools-`uname -r`                                      \
#       linux-tools-common                                          \
#       linux-tools-generic &&                                      \
#     rm -rf /var/lib/apt/lists/*

########################################################################
########################################################################
# Golang
########################################################################

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

########################################################################
########################################################################
# TensorFlow
########################################################################

# Pick up some TF dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential \
        curl \
        libfreetype6-dev \
        libpng12-dev \
        libzmq3-dev \
        pkg-config \
        python \
        python-dev \
        rsync \
        software-properties-common \
        unzip \
        && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN curl -O https://bootstrap.pypa.io/get-pip.py && \
    python get-pip.py && \
    rm get-pip.py

RUN pip --no-cache-dir install \
        ipykernel \
        jupyter \
        matplotlib \
        numpy \
        scipy \
        sklearn \
        Pillow \
        && \
    python -m ipykernel.kernelspec

ENV TENSORFLOW_VERSION 0.12.0rc0

# --- DO NOT EDIT OR DELETE BETWEEN THE LINES --- #
# These lines will be edited automatically by parameterized_docker_build.sh. #
# COPY _PIP_FILE_ /
# RUN pip --no-cache-dir install /_PIP_FILE_
# RUN rm -f /_PIP_FILE_

# Install TensorFlow GPU version.
RUN python2.7 -m pip --no-cache-dir install \
    https://storage.googleapis.com/tensorflow/linux/gpu/tensorflow_gpu-${TENSORFLOW_VERSION}-cp27-none-linux_x86_64.whl
# --- ~ DO NOT EDIT OR DELETE BETWEEN THE LINES --- #


########################################################################
########################################################################
# MXNet
########################################################################

# install the core library
RUN apt-get update && apt-get install -y build-essential git libopenblas-dev libopencv-dev
RUN git clone --recursive https://github.com/dmlc/mxnet/ && cd mxnet && \
    cp make/config.mk . && \
    echo "USE_CUDA=1" >>config.mk && \
    echo "USE_CUDA_PATH=/usr/local/cuda" >>config.mk && \
    echo "USE_CUDNN=1" >>config.mk && \
    echo "USE_BLAS=openblas" >>config.mk && \
    make -j$(nproc) ADD_LDFLAGS=-L/usr/local/cuda/lib64/stubs
ENV LD_LIBRARY_PATH /usr/local/cuda/lib64:$LD_LIBRARY_PATH


########################################################################
########################################################################


RUN mkdir -p /build /src /data

WORKDIR /build
