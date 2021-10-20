# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app
ADD . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY *.env ./

RUN go build -o /api_jwt_gin
EXPOSE 4500

CMD [ "/api_jwt_gin" ]