FROM webgpu/cuda_cudnn:ppc64le_8.0
MAINTAINER Abdul Dakkak <dakkak@illinois.edu>

RUN [ "cross-build-start" ]

# install the core library
RUN apt-get update && apt-get install -y build-essential git libopenblas-dev libopencv-dev
RUN [ "cross-build-end" ]  

WORKDIR /

RUN git clone --recursive https://github.com/dmlc/mxnet
RUN tar -xf /tmp/v0.9.3.tar.gz && \
    mv /mxnet-0.9.3 /mxnet
RUN cd /mxnet && \
    cp make/config.mk . && \
    echo "USE_CUDA=1" >>config.mk && \
    echo "USE_CUDA_PATH=/usr/local/cuda" >>config.mk && \
    echo "USE_CUDNN=1" >>config.mk && \
    echo "USE_BLAS=openblas" >>config.mk && \
    make -j$(nproc) ADD_LDFLAGS=-L/usr/local/cuda/lib64/stubs
ENV LD_LIBRARY_PATH /usr/local/cuda/lib64:$LD_LIBRARY_PATH

RUN [ "cross-build-end" ]  
