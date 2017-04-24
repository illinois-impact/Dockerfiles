FROM webgpu/cuda_cudnn:ppc64le_8.0
MAINTAINER Abdul Dakkak <dakkak@illinois.edu>

RUN [ "cross-build-start" ]

# copy from https://github.com/tensorflow/tensorflow/blob/master/tensorflow/tools/docker/Dockerfile.devel-gpu
# and modifications from https://www.ibm.com/developerworks/community/blogs/fe313521-2e95-46f2-817d-44a4f27eba32/entry/Building_TensorFlow_on_OpenPOWER_Linux_Systems?lang=en

RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential \
        curl \
        git \
        libcurl3-dev \
        libfreetype6-dev \
        libpng12-dev \
        libzmq3-dev \
        pkg-config \
        python-dev \
        rsync \
        software-properties-common \
        unzip \
        zip \
        zlib1g-dev \
        && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ADD https://bootstrap.pypa.io/get-pip.py /tmp/get-pip.py
RUN python /tmp/get-pip.py && \
    rm get-pip.py

# Set up Bazel.

# We need to add a custom PPA to pick up JDK8, since trusty doesn't
# have an openjdk8 backport.  openjdk-r is maintained by a reliable contributor:
# Matthias Klose (https://launchpad.net/~doko).  It will do until
# we either update the base image beyond 14.04 or openjdk-8 is
# finally backported to trusty; see e.g.
#   https://bugs.launchpad.net/trusty-backports/+bug/1368094
RUN echo 'deb http://ppa.launchpad.net/openjdk-r/ppa/ubuntu xenial main' > /etc/apt/sources.list.d/openjdk.list
RUN apt-get update && \
    apt-get install -y --no-install-recommends openjdk-8-jdk

RUN update-alternatives --config java
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-ppc64el

# Running bazel inside a `docker build` command causes trouble, cf:
#   https://github.com/bazelbuild/bazel/issues/134
# The easiest solution is to set up a bazelrc file forcing --batch.
RUN echo "startup --batch" >>/etc/bazel.bazelrc
# Similarly, we need to workaround sandboxing issues:
#   https://github.com/bazelbuild/bazel/issues/418
RUN echo "build --spawn_strategy=standalone --genrule_strategy=standalone" \
    >>/etc/bazel.bazelrc
# Install the most recent bazel release.
ENV BAZEL_VERSION 0.4.5
WORKDIR /
# ADD https://github.com/ibmsoe/bazel/archive/v0.4.5-ppc.tar.gz /bazel
ADD ./v0.4.5-ppc.tar.gz /
RUN ls -l /bazel-0.4.5-ppc
RUN /bazel-0.4.5-ppc/compile.sh

# Download and build TensorFlow.

RUN git clone https://github.com/tensorflow/tensorflow.git && \
    cd tensorflow && \
    git checkout r1.1
WORKDIR /tensorflow

# Configure the build for our CUDA configuration.
ENV CI_BUILD_PYTHON python
ENV LD_LIBRARY_PATH /usr/local/cuda/extras/CUPTI/lib64:$LD_LIBRARY_PATH
ENV TF_NEED_CUDA 1
ENV TF_CUDA_COMPUTE_CAPABILITIES=3.0,3.5,5.2,6.0,6.1

RUN tensorflow/tools/ci_build/builds/configured GPU \
    bazel build -c opt --config=cuda tensorflow/tools/pip_package:build_pip_package && \
    bazel-bin/tensorflow/tools/pip_package/build_pip_package /tmp/pip && \
    pip --no-cache-dir install --upgrade /tmp/pip/tensorflow-*.whl && \
    rm -rf /tmp/pip && \
    rm -rf /root/.cache
# Clean up pip wheel and Bazel cache when done.

RUN [ "cross-build-end" ]  