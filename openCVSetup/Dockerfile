FROM python:2.7.12
RUN apt-get update && apt-get install -y build-essential cmake git\
 libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev
RUN apt-get update && apt-get install -y python-dev \
 python-numpy libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev \
libjasper-dev libdc1394-22-dev unzip
RUN wget -O opencv-2.4.13.zip http://sourceforge.net/projects/opencvlibrary/files/opencv-unix/2.4.13/opencv-2.4.13.zip/download
RUN unzip opencv-2.4.13.zip
RUN mv opencv-2.4.13 opencv
WORKDIR /opencv/build

RUN virtualenv ~/cv/
RUN . ~/cv/bin/activate

RUN export PYTHON2_EXECUTRABLE=/usr/local/bin/python
RUN export PYTHON_INCLUDE_DIR=/usr/local/include/python2.7
RUN export PYTHON_INCLUDE_DIR2=/usr/include/x86_64-linux-gnu/python2.7
RUN export PYTHON_LIBRARY=/usr/lib/x86_64-linux-gnu/libpython2.7.so
RUN export PYTHON_NUMPY_INCLUDE_DIRS=/usr/lib/python2.7/dist-packages/numpy/core/include/

RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_PREFIX=/usr/local \
-D BUILD_NEW_PYTHON_SUPPORT=ON -D INSTALL_PYTHON_EXAMPLES=ON -D BUILD_EXAMPLES=ON \
/opencv

RUN make -j7 # runs 7 jobs in parallel.
RUN make install
RUN ldconfig
WORKDIR ~/cv/lib/python2.7/site-packages/
RUN ln -s /usr/local/lib/python2.7/site-packages/cv2.so cv2.so
RUN ln -s /usr/local/lib/python2.7/site-packages/cv.py cv.py
COPY ./bash_profile ~/.bash_profile
