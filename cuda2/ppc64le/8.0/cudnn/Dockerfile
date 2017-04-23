FROM webgpu/cuda_devel:ppc64le_8.0
MAINTAINER Abdul Dakkak <dakkak@illinois.edu>

RUN [ "cross-build-start" ]

ENV CUDNN_VERSION 6
LABEL com.nvidia.cudnn.version="6"

ADD http://developer.download.nvidia.com/compute/redist/cudnn/v6.0/cudnn-8.0-linux-ppc64le-v6.0.tgz /usr/local 
RUN ldconfig

RUN ls -l /usr/local

RUN [ "cross-build-end" ]  