# SpyCam007
A project for CMPT 436. We're going to do motion capturing with a pi, upload it to a server, and serve it on a web app. 

#Installation Instructions
## SpyCam Software
The spy cam runs as a python script, but in order to get it to work, OpenCV, an open source computer vision package needs to be installed
detailed instructions on how to install OpenCV on a Pi can be found [here](http://www.pyimagesearch.com/2016/04/18/install-guide-raspberry-pi-3-raspbian-jessie-opencv-3/)
This script sends images via POST to a webserver, if you want to run your own version, the code will have to be altered to send to your own server.
```
 #line 117#
 r = requests.post('http://laforgesplayground.appspot.com/pictures/', files={'picture': open(imgName, 'rb')})
```
Once open cv is installed and the requests are being sent to your server (installation described below) you can start the camera
by simply calling ```python detector.py``` the script will use your first webcam available, usb or built in on a laptop.

The script will take images at 10 second interval provided there is motion.




