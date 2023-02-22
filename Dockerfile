FROM golang:1.18-alpine

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod /build
ADD go.sum /build
RUN go mod download

COPY . /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main main.go 

ENTRYPOINT ["/main"]

