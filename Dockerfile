# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ADD . .
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY *.env ./

RUN go build -o /go_jwt
EXPOSE 4500

CMD [ "/go_jwt" ]