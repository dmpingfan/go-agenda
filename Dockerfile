FROM golang:1.8

WORKDIR /
ADD . "$GOPATH/src/github.com/painterdrown/go-agenda"
RUN cd "$GOPATH/src/github.com/painterdrown/go-agenda/cli" && go get -u && go install
RUN cd "$GOPATH/src/github.com/painterdrown/go-agenda/service" && go get -u && go install
EXPOSE 3000
VOLUME ["/data"]