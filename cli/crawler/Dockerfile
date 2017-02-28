FROM sfmovies-core:latest

WORKDIR /opt/sfmovies/cli/crawler
RUN go build client.go
RUN go build server.go

ENTRYPOINT /opt/sfmovies/cli/crawler/server run | tee /var/log/sfmovies/crawler.log
EXPOSE 8080
