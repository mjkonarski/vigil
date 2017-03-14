FROM golang:1.8

RUN mkdir -p /go/src/vigil
WORKDIR /go/src/vigil
COPY . /go/src/vigil

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]