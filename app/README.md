# Setting up appengine golang
Simple quickstart instructions: https://cloud.google.com/appengine/docs/go/quickstart

Download the go appengine_sdk here: https://cloud.google.com/appengine/docs/go/download

add go appengine_sdk to your path: 
in .bash_profile: PATH=$PATH:(path to sdk root folder)

type

`goapp`. 

If it worked, you installed it correctly.


# Using the Server
## Installing dependencies

For online dependencies, run 

`go get -d -t -v . && go get -d -t -v github.com/SpyCam007`

`go install . 

For internal dependencies, run: 

`go install github.com/SpyCam007'


## Running the project

run project: 

`goapp serve ./`

It should now be serving on port:8080.

deployment: 

`appcfg.py -A laforgesplayground -V v1 update ./`


#Setup Go

On the command line, you should configure GOROOT and GOPATH env variables.
I have a script that works for me (on mac) 

`./setup_go.`

