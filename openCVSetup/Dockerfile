# KEEP UBUNTU OR DEBIAN UP TO DATE
FROM python:2.7.12
# INSTALL THE DEPENDENCIES

# Build tools:
RUN apt-get install -y build-essential cmake

# GUI (if you want to use GTK instead of Qt, replace 'qt5-default' with 'libgtkglext1-dev' and remove '-DWITH_QT=ON' option in CMake):
RUN apt-get install -y qt5-default libvtk6-dev

# Media I/O:
RUN apt-get install -y zlib1g-dev libjpeg-dev libwebp-dev libpng-dev libtiff5-dev libjasper-dev libopenexr-dev libgdal-dev

# Video I/O:
RUN apt-get install -y libdc1394-22-dev libavcodec-dev libavformat-dev libswscale-dev libtheora-dev libvorbis-dev libxvidcore-dev libx264-dev yasm libopencore-amrnb-dev libopencore-amrwb-dev libv4l-dev libxine2-dev

# Parallelism and linear algebra libraries:
RUN apt-get install -y libtbb-dev libeigen3-dev

# Python:
RUN apt-get install -y python-dev python-tk python-numpy python3-dev python3-tk python3-numpy

# Java:
RUN apt-get install -y ant default-jdk

# Documentation:
RUN apt-get install -y doxygen


# INSTALL THE LIBRARY (YOU CAN CHANGE '3.1.0' FOR THE LAST STABLE VERSION)

RUN apt-get install -y unzip wget
RUN wget https://github.com/Itseez/opencv/archive/3.1.0.zip
RUN unzip 3.1.0.zip
RUN rm 3.1.0.zip
RUN mv opencv-3.1.0 OpenCV
RUN cd OpenCV
RUN mkdir build
WORKDIR /OpenCV/build
RUN cmake -DWITH_QT=ON -DWITH_OPENGL=ON -DFORCE_VTK=ON -DWITH_TBB=ON -DWITH_GDAL=ON -DWITH_XINE=ON -DBUILD_EXAMPLES=ON ..
RUN make -j4
RUN make install
RUN ldconfig
