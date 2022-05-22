FROM golang:latest

WORKDIR /app

MAINTAINER YourName

ADD . /app/

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod verify

EXPOSE 8080

CMD ["go", "build", "main.go"]