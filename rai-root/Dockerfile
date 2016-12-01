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
# Tensorflow
########################################################################


RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential \
        curl \
        git \
        libcurl3-dev \
        libfreetype6-dev \
        libpng12-dev \
        libzmq3-dev \
        pkg-config \
        python \
        python-dev \
        rsync \
        software-properties-common \
        unzip \
        zip \
        zlib1g-dev \
        && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN curl -fSsL -O https://bootstrap.pypa.io/get-pip.py && \
    python get-pip.py && \
    rm get-pip.py

RUN pip --no-cache-dir install \
        ipykernel \
        jupyter \
        matplotlib \
        numpy \
        scipy \
        sklearn \
        && \
    python -m ipykernel.kernelspec

# Set up Bazel.

# We need to add a custom PPA to pick up JDK8, since trusty doesn't
# have an openjdk8 backport.  openjdk-r is maintained by a reliable contributor:
# Matthias Klose (https://launchpad.net/~doko).  It will do until
# we either update the base image beyond 14.04 or openjdk-8 is
# finally backported to trusty; see e.g.
#   https://bugs.launchpad.net/trusty-backports/+bug/1368094
RUN add-apt-repository -y ppa:openjdk-r/ppa && \
    apt-get update && \
    apt-get install -y --no-install-recommends openjdk-8-jdk openjdk-8-jre-headless && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Running bazel inside a `docker build` command causes trouble, cf:
#   https://github.com/bazelbuild/bazel/issues/134
# The easiest solution is to set up a bazelrc file forcing --batch.
RUN echo "startup --batch" >>/root/.bazelrc
# Similarly, we need to workaround sandboxing issues:
#   https://github.com/bazelbuild/bazel/issues/418
RUN echo "build --spawn_strategy=standalone --genrule_strategy=standalone" \
    >>/root/.bazelrc
ENV BAZELRC /root/.bazelrc
# Install the most recent bazel release.
ENV BAZEL_VERSION 0.3.2
WORKDIR /
RUN mkdir /bazel && \
    cd /bazel && \
    curl -fSsL -O https://github.com/bazelbuild/bazel/releases/download/$BAZEL_VERSION/bazel-$BAZEL_VERSION-installer-linux-x86_64.sh && \
    curl -fSsL -o /bazel/LICENSE.txt https://raw.githubusercontent.com/bazelbuild/bazel/master/LICENSE.txt && \
    chmod +x bazel-*.sh && \
    ./bazel-$BAZEL_VERSION-installer-linux-x86_64.sh && \
    cd / && \
    rm -f /bazel/bazel-$BAZEL_VERSION-installer-linux-x86_64.sh

# Download and build TensorFlow.

RUN git clone https://github.com/tensorflow/tensorflow.git && \
    cd tensorflow && \
    git checkout r0.12
WORKDIR /tensorflow

# Configure the build for our CUDA configuration.
ENV LD_LIBRARY_PATH /usr/local/cuda/extras/CUPTI/lib64:$LD_LIBRARY_PATH
ENV TF_NEED_CUDA 1
ENV TF_CUDA_COMPUTE_CAPABILITIES=3.0,3.5,5.2

RUN tensorflow/tools/ci_build/builds/configured GPU \
    bazel build -c opt --config=cuda tensorflow/tools/pip_package:build_pip_package && \
    bazel-bin/tensorflow/tools/pip_package/build_pip_package /tmp/pip && \
    pip install --upgrade /tmp/pip/tensorflow-*.whl && \
    rm -rf /tmp/pip && \
    rm -rf /root/.cache
# Clean up pip wheel and Bazel cache when done.

########################################################################
########################################################################
# Torch
########################################################################

# Install git, apt-add-repository and dependencies for iTorch
RUN apt-get update && apt-get install -y \
  git \
  software-properties-common \
  ipython3 \
  libssl-dev \
  libzmq3-dev \
  python-zmq \
  python-pip

# Run Torch7 installation scripts (dependencies only)
RUN git clone https://github.com/torch/distro.git /root/torch --recursive && cd /root/torch && \
  bash install-deps


# Run Torch7 installation scripts
RUN cd /root/torch && \
# Run without nvcc to prevent timeouts
  sed -i 's/path_to_nvcc=$(which nvcc)/path_to_nvcc=$(which no_nvcc)/g' install.sh && \
  sed -i 's,path_to_nvcc=/usr/local/cuda/bin/nvcc,path_to_nvcc=,g' install.sh && \
  ./install.sh

# Export environment variables manually
ENV LUA_PATH='/root/.luarocks/share/lua/5.1/?.lua;/root/.luarocks/share/lua/5.1/?/init.lua;/root/torch/install/share/lua/5.1/?.lua;/root/torch/install/share/lua/5.1/?/init.lua;./?.lua;/root/torch/install/share/luajit-2.1.0-beta1/?.lua;/usr/local/share/lua/5.1/?.lua;/usr/local/share/lua/5.1/?/init.lua'
ENV LUA_CPATH='/root/.luarocks/lib/lua/5.1/?.so;/root/torch/install/lib/lua/5.1/?.so;./?.so;/usr/local/lib/lua/5.1/?.so;/usr/local/lib/lua/5.1/loadall.so'
ENV PATH=/root/torch/install/bin:$PATH
ENV LD_LIBRARY_PATH=/root/torch/install/lib:$LD_LIBRARY_PATH
ENV DYLD_LIBRARY_PATH=/root/torch/install/lib:$DYLD_LIBRARY_PATH
ENV LUA_CPATH='/root/torch/install/lib/?.so;'$LUA_CPATH


# Restore Torch7 installation script
RUN sed -i 's/path_to_nvcc=$(which no_nvcc)/path_to_nvcc=$(which nvcc)/g' install.sh

# Install CUDA libraries
RUN luarocks install cutorch && \
  luarocks install cunn && \
  luarocks install cudnn

########################################################################
########################################################################
# MXNet
########################################################################

# Install git, wget and other dependencies
RUN apt-get update && apt-get install -y \
  git \
  libopenblas-dev \
  libopencv-dev \
  python-dev \
  python-numpy \
  python-setuptools \
  wget

# Clone MXNet repo and move into it
RUN cd /root && git clone --recursive https://github.com/dmlc/mxnet && cd mxnet && \
# Copy config.mk
  cp make/config.mk config.mk && \
# Set OpenBLAS
  sed -i 's/USE_BLAS = atlas/USE_BLAS = openblas/g' config.mk && \
# Set CUDA flag
  sed -i 's/USE_CUDA = 0/USE_CUDA = 1/g' config.mk && \
  sed -i 's/USE_CUDA_PATH = NONE/USE_CUDA_PATH = \/usr\/local\/cuda/g' config.mk && \
# Set cuDNN flag
  sed -i 's/USE_CUDNN = 0/USE_CUDNN = 1/g' config.mk && \
# Make
  make -j"$(nproc)"

# Install Python package
RUN cd /root/mxnet/python && python setup.py install


########################################################################
########################################################################


RUN mkdir -p /build /src /data

WORKDIR /build
