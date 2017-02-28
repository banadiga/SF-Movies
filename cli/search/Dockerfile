FROM sfmovies-core:latest

WORKDIR /opt/sfmovies/cli/search
RUN go build client.go
RUN go build server.go

ENTRYPOINT /opt/sfmovies/cli/search/server run | tee /var/log/sfmovies/search.log
EXPOSE 8080
