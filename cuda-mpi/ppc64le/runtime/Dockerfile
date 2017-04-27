FROM webgpu/cuda:ppc64le-8.0-runtime
MAINTAINER Carl Pearson <pearson@illinois.edu>

RUN apt-get update && apt-get -y --no-install-recommends install wget cmake

RUN apt-get update && apt-get install -y --no-install-recommends \
        mpich && \
    rm -rf /var/lib/apt/lists/*

# ENV PATH /usr/local/nvidia/bin:/usr/local/cuda/bin:${PATH}
# ENV LD_LIBRARY_PATH /usr/local/nvidia/lib:/usr/local/nvidia/lib64:${LD_LIBRARY_PATH}
