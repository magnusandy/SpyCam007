FROM python:2.7.12
RUN apt-get update && apt-get install -y build-essential cmake git\
 libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev \
 python-dev python-numpy libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev \
libjasper-dev libdc1394-22-dev

WORKDIR /openCV/
COPY ./opencv-2.4.13 ./opencv
WORKDIR /openCV/opencv/build

RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_PREFIX=/usr/local /openCV/opencv

RUN export PYTHON2_EXECUTRABLE=/usr/local/bin/python
RUN export PYTHON_INCLUDE_DIR=/usr/local/include/python2.7
RUN export PYTHON_INCLUDE_DIR2=/usr/include/x86_64-linux-gnu/python2.7
RUN export PYTHON_LIBRARY=/usr/lib/x86_64-linux-gnu/libpython2.7.so
RUN export PYTHON_NUMPY_INCLUDE_DIRS=/usr/lib/python2.7/dist-packages/numpy/core/include/

RUN make -j7 # runs 7 jobs in parallel.
RUN make install 

COPY ./bash_profile ~/.bash_profile
