FROM webgpu/cuda_devel:ppc64le_8.0
MAINTAINER Abdul Dakkak <dakkak@illinois.edu>

RUN [ "cross-build-start" ]

# Set CUDA_ROOT
ENV CUDA_ROOT /usr/local/cuda/bin

# Follow directions found in
# http://docs.opencv.org/3.2.0/d6/d15/tutorial_building_tegra_cuda.html

# Install dependencies
RUN \ 
    apt-get update && apt-get upgrade -y && \
    apt-get install \
        build-essential \
        cmake \
        git \
        libglew-dev \
        libtiff-dev \
        zlib1g-dev \
        libjpeg-dev \
        libpng12-dev \
        libjasper-dev \
        libavcodec-dev \
        libavformat-dev \
        libavutil-dev \
        libpostproc-dev \
        libswscale-dev \
        libeigen3-dev \
        libtbb-dev \
        libgtk2.0-dev \
        pkg-config \
        libatlas-base-dev \
        gfortran \
        libhdf5-dev \
        python-dev \
        python-numpy \
        python-py \
        python-pytest \
        python-pip -y && \
    apt-get autoclean && apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


# Downloads and configure OpenCV
ADD https://github.com/opencv/opencv/archive/3.2.0.tar.gz /root/opencv
ADD https://github.com/opencv/opencv_contrib/archive/3.2.0.tar.gz /root/opencv_contrib

RUN mv /root/opencv-3.2.0 /root/opencv && \
    mv /root/opencv_contrib-3.2.0 /root/opencv_contrib

RUN cd /root/opencv && \
    mkdir build && \
    cd build && \
    cmake \
        -DCMAKE_BUILD_TYPE=Release \
        -DCMAKE_INSTALL_PREFIX=/usr \
        -DBUILD_PNG=OFF \
        -DBUILD_TIFF=OFF \
        -DBUILD_TBB=OFF \
        -DBUILD_JPEG=OFF \
        -DBUILD_JASPER=OFF \
        -DBUILD_ZLIB=OFF \
        -DBUILD_EXAMPLES=OFF \
        -DBUILD_opencv_java=OFF \
        -DBUILD_opencv_python2=ON \
        -DBUILD_opencv_python3=OFF \
        -DWITH_OPENCL=OFF \
        -DWITH_OPENMP=OFF \
        -DWITH_FFMPEG=ON \
        -DWITH_GSTREAMER=OFF \
        -DWITH_GSTREAMER_0_10=OFF \
        -DWITH_CUDA=ON \
        -DWITH_GTK=ON \
        -DWITH_VTK=OFF \
        -DWITH_TBB=ON \
        -DWITH_1394=OFF \
        -DWITH_OPENEXR=OFF \
        -DCUDA_TOOLKIT_ROOT_DIR=/usr/local/cuda \
        -DCUDA_ARCH_BIN='3.0 3.5 5.0 6.0 6.2' \
        -DCUDA_ARCH_PTX="" \
        -DINSTALL_C_EXAMPLES=ON \
        -DINSTALL_TESTS=OFF \
        -DOPENCV_TEST_DATA_PATH=../opencv_extra/testdata \
        .. &&\
	cd ~/opencv/build && \
    make -j $(nproc) VERBOSE=1 && \
    make install && \
    ldconfig && \

    cp ~/opencv/build/lib/cv2.so /usr/local/lib/python2.7/dist-packages/


RUN [ "cross-build-end" ]  