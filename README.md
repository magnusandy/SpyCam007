# SpyCam007
A project for CMPT 436. We're going to do motion capturing with a pi, upload it to a server, and serve it on a web app. 

[Detailed Presentation and Description](https://drive.google.com/open?id=1knmU_6cxQ53bzwLr8JVPvMZ8jszvsUK12UYOAXiN_m4)

# Screenshots
![img 1](http://i.imgur.com/BhFkkzt.png?1)
![img 1](http://i.imgur.com/yg9wc26.png)
![img 1](http://i.imgur.com/wDM8rsF.jpg?1)
# Installation Instructions
## SpyCam Software
The spy cam runs as a python script, but in order to get it to work, OpenCV, an open source computer vision package needs to be installed
detailed instructions on how to install OpenCV on a Pi can be found [here](http://www.pyimagesearch.com/2016/04/18/install-guide-raspberry-pi-3-raspbian-jessie-opencv-3/)

other packages are necessary and can be installed via pip
```
pip install requests
```
This script sends images via POST to a webserver, if you want to run your own version, the code will have to be altered to send to your own server.
```
 #line 117#
 r = requests.post('http://laforgesplayground.appspot.com/pictures/', files={'picture': open(imgName, 'rb')})
```
Once open cv is installed and the requests are being sent to your server (installation described below) you can start the camera
by simply calling ```python detector.py``` the script will use your first webcam available, usb or built in on a laptop.

The script will take images at 10 second interval provided there is motion.


## Appengine golang
### Simple quickstart instructions: https://cloud.google.com/appengine/docs/go/quickstart

### Download the go appengine_sdk here: https://cloud.google.com/appengine/docs/go/download

### add go appengine_sdk to your path 
in .bash_profile: 
``` bash
PATH=$PATH:(path to sdk root folder)
```

# Setting up goroot / gopath
either use the command
```bash 
. setup_go
```
in the server directory if you downloaded go_appengine to usr/local/go_appengine, otherwise use:
```bash
export $GOROOT={path to go_appengine/goroot}
export $GOPATH={path to go_appengine/gopath}
```
# Install the dependencies
type: 
```bash
go get -u "google.golang.org/appengine/urlfetch"
go get -u "cloud.google.com/go/storage"
```

# Running the dev app server
Run 
```bash 
goapp server
```
In the server directory. It should now be serving on port:8080.

# Deploying the app
Type
```bash
goapp deploy -application laforgesplayground --version v1 app.yaml
```

This will upload the changes to laforgesplayground.appspot.com.

