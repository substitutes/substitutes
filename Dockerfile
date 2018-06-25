FROM golang:1.8

WORKDIR /go/src/substitutes
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV GIN_MODE release 

CMD ["substitutes"]
