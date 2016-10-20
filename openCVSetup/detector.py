# import the necessary packages
import argparse
import datetime
#import imutils #ignoring for now
import time
import cv2

firstFrame = None;
imgScale = 1; #change the image size
minFindArea = 50; #min size of find, removes small pieces of movement
#used by the gaussian blur function, these are the sizes for the area that the function will blur over in pixels
blurKernalSizeX = 5;
blurKernalSizeY = 5;
blurKernalSigmaX = 0;
#threshhold Values
greyscaleThreshholdValue = 25 # on a scale of 0-255 (white to black) the threshhold is the cutoff for which pixels are kept during the threshhold function if below they are kept(and turned white) and above are turned blackhamster124
greyscaleBlack = 255

#number of passes of dialation, higher means larger merging
dilationMultiplier = 2;
#starts the videocam on the default channel 0
camera = cv2.VideoCapture(0); #starts the videocam on the default channel 0
print camera.isOpened()
time.sleep(20); #sleep long enough for the camera to boot up and what not

#main loop of the program, will continue
while True:
    text = "Empty"
    #read in the first frame from the video stream
    (frameReadCorrectly, readFrame) = camera.read();
    if not frameReadCorrectly:
        break
    #resize frame to half size fx and fy are size scales
    readFrame = cv2.resize(readFrame, (500,500))
    cv2.imshow('Initial Frame before processing', readFrame)
    #converts the frame to black and white and preforms a blur on the image
    #smoothing out the image and reducing "noise" making processing easier later
    blackWhiteFrame = cv2.cvtColor(readFrame, cv2.COLOR_BGR2GRAY);
    blackWhiteFrame = cv2.GaussianBlur(blackWhiteFrame, (blurKernalSizeX, blurKernalSizeY), blurKernalSigmaX);
    if firstFrame is None:
        firstFrame = blackWhiteFrame;
        continue
    cv2.imshow('firstFrame', firstFrame)
    cv2.imshow('blackWhite', blackWhiteFrame)

    #take the difference betwen initial frame and current Frame
    frameDelta = cv2.absdiff(firstFrame, blackWhiteFrame)
    #threshhold converts any pixels below the greyscaleThreshholdValue to white and above to black
    ret, threshholdFrame = cv2.threshold(frameDelta, greyscaleThreshholdValue, greyscaleBlack, cv2.THRESH_BINARY)
    cv2.imshow('frameDelta', frameDelta)
    cv2.imshow('thresh', threshholdFrame)

    #used to clean up and connect the pieces of movement that are close together into one chunk
    dilatedFrame = cv2.dilate(threshholdFrame, None, iterations=dilationMultiplier)
    cv2.imshow('dilated', dilatedFrame)

    #finds all the seperate shapes in white in the image (all the pieces of movement)
    (_, contours, _) = cv2.findContours(dilatedFrame.copy(), cv2.RETR_EXTERNAL,
		cv2.CHAIN_APPROX_SIMPLE)
    for c in contours:
        if cv2.contourArea(c) < minFindArea:
            continue
        (x, y, w, h) = cv2.boundingRect(c)
        cv2.rectangle(readFrame, (x, y), (x + w, y + h), (0, 255, 0), 2)
        text = "Occupied"

    cv2.putText(readFrame, "Room Status: {}".format(text), (10, 20), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 0, 255), 2)
    cv2.putText(readFrame, datetime.datetime.now().strftime("%A %d %B %Y %I:%M:%S%p"), (10, readFrame.shape[0] - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.35, (0, 0, 255))
    cv2.imshow('final', readFrame)
    if cv2.waitKey(1) == 27:
            break  # esc to quit
cv2.destroyAllWindows()
    
