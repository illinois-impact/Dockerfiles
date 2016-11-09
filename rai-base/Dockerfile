FROM nvidia/cuda:8.0-devel-ubuntu16.04

MAINTAINER Abdul Dakkak "dakkak@illinois.edu"

# Set one or more individual labels
LABEL com.webgpu.rai.version="0.0.1"
LABEL vendor="RAI Base Image"

#Build Essentials
RUN apt-get update && apt-get install --no-install-recommends -y  \
      software-properties-common                                  \
      python-software-properties &&                               \
    add-apt-repository ppa:ubuntu-toolchain-r/test &&             \
    apt-get update && apt-get install --no-install-recommends -y  \
      build-essential                                             \
      gcc                                                         \
      clang                                                       \
      ntp                                                         \
      bzip2                                                       \
      autoconf                                                    \
      automake                                                    \
      make                                                        \
      tar                                                         \
      unzip                                                       \
      zlib1g-dbg                                                  \
      wget                                                        \
      git                                                         \
      file                                                        \
      libhdf5-dev                                                 \
      libhdf5-10                                                  \
      libhdf5-10-dbg                                              \
      libhdf5-openmpi-10                                          \
      libhdf5-openmpi-10-dbg                                      \
      libhdf5-openmpi-dev                                         \
      libhdf5-cpp-11                                              \
      libhdf5-cpp-11-dbg                                          \
      libopenmpi-dev                                              \
      cmake &&                                                    \
    rm -rf /var/lib/apt/lists/*

# RUN ntpdate pool.ntp.org

#Setup User
ENV USERNAME raiuser
RUN useradd -ms /bin/bash ${USERNAME}
ENV USERDIR /home/${USERNAME}

ENV HOME ${USERDIR}
ENV SRCDIR ${HOME}/src
ENV DATADIR ${HOME}/data
ENV BUILDDIR ${HOME}/build

RUN chown -R ${USERNAME} ${HOME}

WORKDIR ${BUILDDIR}
USER ${USERNAME}
