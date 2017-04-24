FROM multiarch/ubuntu-core:ppc64el-xenial
MAINTAINER Abdul Dakkak <dakkak@illinois.edu>

LABEL com.nvidia.volumes.needed="rai-cuda"

ENV ARCH ppc64le
ENV QEMU_EXECVE 1

COPY . /usr/bin

RUN [ "/usr/bin/qemu-ppc64le-static", "/bin/sh", "-c", "ln -s /usr/bin/cross_build /usr/bin/cross-build-start; ln -s /usr/bin/cross_build /usr/bin/cross-build-end; cp /bin/sh /bin/sh.real" ]

RUN [ "cross-build-start" ]

RUN apt-get update && apt-get -y --no-install-recommends install wget cmake git

WORKDIR /tmp
RUN CUDA_REPO_URL="http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1604/ppc64el/cuda-repo-ubuntu1604_8.0.61-1_ppc64el.deb" && \
    wget ${CUDA_REPO_URL} && dpkg --install *.deb && rm -rf *.deb
    

ENV CUDA_VERSION 8.0
LABEL com.nvidia.cuda.version="8.0"

ENV CUDA_PKG_VERSION 8-0=8.0.61-1
RUN apt-get update && apt-get install -y --no-install-recommends \
        cuda-nvrtc-$CUDA_PKG_VERSION \
        cuda-nvgraph-$CUDA_PKG_VERSION \
        cuda-cusolver-$CUDA_PKG_VERSION \
        cuda-cublas-$CUDA_PKG_VERSION \
        cuda-cufft-$CUDA_PKG_VERSION \
        cuda-curand-$CUDA_PKG_VERSION \
        cuda-cusparse-$CUDA_PKG_VERSION \
        cuda-npp-$CUDA_PKG_VERSION \
        cuda-cudart-$CUDA_PKG_VERSION && \
    ln -s cuda-$CUDA_VERSION /usr/local/cuda && \
    rm -rf /var/lib/apt/lists/*

RUN echo "/usr/local/cuda/lib" >> /etc/ld.so.conf.d/cuda.conf && \
    echo "/usr/local/cuda/lib64" >> /etc/ld.so.conf.d/cuda.conf && \
    ldconfig

RUN echo "/usr/local/nvidia/lib" >> /etc/ld.so.conf.d/nvidia.conf && \
    echo "/usr/local/nvidia/lib64" >> /etc/ld.so.conf.d/nvidia.conf

ENV PATH /usr/local/nvidia/bin:/usr/local/cuda/bin:${PATH}
ENV LD_LIBRARY_PATH /usr/local/nvidia/lib:/usr/local/nvidia/lib64:${LD_LIBRARY_PATH}


RUN apt-get update && apt-get -y --no-install-recommends install wget cmake gcc-4.9 g++-4.9 && \
     update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.9 60 --slave /usr/bin/g++ g++ /usr/bin/g++-4.9 && \
     update-alternatives --config gcc

RUN [ "cross-build-end" ]  
