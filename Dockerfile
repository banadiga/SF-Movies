FROM golang

RUN mkdir /opt/sfmovies
ADD . /opt/sfmovies
WORKDIR /opt/sfmovies

RUN go get -d ./...

