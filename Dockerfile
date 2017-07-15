FROM golang:1.8

RUN mkdir -p /go/src/vigil
WORKDIR /go/src/vigil
COPY . /go/src/vigil

RUN go-wrapper download
RUN go-wrapper install

RUN go get -u github.com/pressly/goose/cmd/goose

CMD ["/bin/bash", "run.sh"]
