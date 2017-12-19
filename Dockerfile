FROM golang:latest

WORKDIR /
ADD . "$GOPATH/src/github.com/painterdrown/go-agenda"
RUN cd "$GOPATH/src/github.com/painterdrown/go-agenda/cli" && go get -u && go install
RUN cd "$GOPATH/src/github.com/painterdrown/go-agenda/service" && go get -u && go install
EXPOSE 8080
VOLUME ["/data"]