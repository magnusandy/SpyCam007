FROM python:2.7.12
RUN apt-get update && apt-get install -y build-essential cmake git\
 libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev
 python-dev python-numpy libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libjasper-dev libdc1394-22-dev
WORKDIR /openCV/
COPY ./opencv-2.4.13 ./
