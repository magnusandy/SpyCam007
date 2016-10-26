# import the necessary packages
import argparse
import datetime
#import imutils #ignoring for now
import time
import cv2

threatText = "Threat Detected!"
secureText = "Area Secure"
imageType = ".png"
averageBGFrame = None;
imgScale = 0.5; #change the image size
minFindArea = 100000; #min area of find, removes small pieces of movement
#used by the gaussian blur function, these are the sizes for the area that the function will blur over in pixels
blurKernalSizeX = 11;
blurKernalSizeY = 11;
blurKernalSigmaX = 0;
#threshhold Values
greyscaleThreshholdValue = 25 # on a scale of 0-255 (white to black) the threshhold is the cutoff for which pixels are kept during the threshhold function if below they are kept(and turned white) and above are turned black
greyscaleBlack = 255

accumulateSpeed = 0.01; #speed at which accumulator forgets old values

#number of passes of dialation, higher means larger merging
dilationMultiplier = 4;

lastCaptureTime = time.time();
captureInterval = 10#seconds
#starts the videocam on the default channel 0
camera = cv2.VideoCapture(0); #starts the videocam on the default channel 0
 #sleep long enough for the camera to boot up and what not
time.sleep(2)
#main loop of the program, will continue
while True:
    text = secureText
    #read in the first frame from the video stream
    (frameReadCorrectly, readFrame) = camera.read();
    if not frameReadCorrectly:
        break
    #resize frame to half size fx and fy are size scales
    #readFrame = cv2.resize(readFrame, (0,0), fx=imgScale, fy=imgScale)
    #cv2.imshow('Initial Frame before processing', readFrame)
    #converts the frame to black and white and preforms a blur on the image
    #smoothing out the image and reducing "noise" making processing easier later
    blackWhiteFrame = cv2.cvtColor(readFrame, cv2.COLOR_BGR2GRAY);
    blackWhiteFrame = cv2.GaussianBlur(blackWhiteFrame, (blurKernalSizeX, blurKernalSizeY), blurKernalSigmaX);
    #This initializes the background, which is an average of the frames before it
    #instead of comparing to a static background, this can change over time
    if averageBGFrame is None:
        averageBGFrame = blackWhiteFrame.copy().astype("float")
        continue
    #cv2.imshow('averageBGFrame', averageBGFrame)
    #cv2.imshow('blackWhite', blackWhiteFrame)

    #adds blackWhiteFrame image onto averageBGFrame as a weighted average, the higher the speed the faster the weight will shift from older values
    cv2.accumulateWeighted(blackWhiteFrame, averageBGFrame, accumulateSpeed)

    #take the difference betwen initial frame and current Frame
    frameDelta = cv2.absdiff(blackWhiteFrame, cv2.convertScaleAbs(averageBGFrame))
    #threshhold converts any pixels below the greyscaleThreshholdValue to white and above to black
    ret, threshholdFrame = cv2.threshold(frameDelta, greyscaleThreshholdValue, greyscaleBlack, cv2.THRESH_BINARY)
    #cv2.imshow('frameDelta', frameDelta)
    #cv2.imshow('thresh', threshholdFrame)

    #used to clean up and connect the pieces of movement that are close together into one chunk
    dilatedFrame = cv2.dilate(threshholdFrame, None, iterations=dilationMultiplier)
    #cv2.imshow('dilated', dilatedFrame)

    #finds all the seperate shapes in white in the image (all the pieces of movement)
    (_, contours, _) = cv2.findContours(dilatedFrame.copy(), cv2.RETR_EXTERNAL,
		cv2.CHAIN_APPROX_SIMPLE)

    #This next block of code finds the largest contourArea and draws a box around it on the screen
    biggestAreaFound = 0;
    biggestContour = None;
    for c in contours:
        cArea = cv2.contourArea(c);
        #contours must be larger than a specified size to weed out small pieces of movement
        if cArea > minFindArea:
            #keeps only the biggest contours
            if cArea > biggestAreaFound:
                biggestAreaFound = cArea;
                biggestContour = c;
                text = threatText

    #create the bounding rectangle for the largest contour
    if biggestContour is not None:
        (x, y, w, h) = cv2.boundingRect(biggestContour)
        cv2.rectangle(readFrame, (x, y), (x + w, y + h), (0, 255, 0), 2)

    ##draw the bounding rectangle on the image as well as a  timestamp and status
    cv2.putText(readFrame, text, (10, 20), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 0, 255), 2)
    currentTime = datetime.datetime.now();
    currentTimeString = currentTime.strftime("%A-%d-%B-%Y-%I:%M:%S%p")
    ct = time.time()
    cv2.putText(readFrame, currentTimeString, (10, readFrame.shape[0] - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.35, (0, 0, 255))
    #cv2.imshow('final', readFrame)
    if ((ct-lastCaptureTime) > captureInterval) and text == threatText:
        lastCaptureTime = ct
        imgName = currentTimeString+imageType
        print "saving image: "+imgName
        cv2.imwrite(imgName, readFrame)
    if cv2.waitKey(1) == 27:
            break  # esc to quit
cv2.destroyAllWindows()
