FROM golang:1.16.4-buster as golang

RUN go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
RUN cp $GOPATH/bin/* /usr/local/bin/

FROM namely/protoc-all

WORKDIR /app
COPY --from=golang /usr/local/bin/* /usr/local/bin/
COPY . /app/
