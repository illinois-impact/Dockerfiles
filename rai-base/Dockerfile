FROM webgpu/rai:root

MAINTAINER Abdul Dakkak "dakkak@illinois.edu"

# Set one or more individual labels
LABEL com.webgpu.rai.version="0.0.1"
LABEL vendor="RAI Base Image"

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
