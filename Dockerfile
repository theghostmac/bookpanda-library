FROM golang:latest

WORKDIR /app

MAINTAINER YourName

ADD . /app/

EXPOSE 8080

CMD ["go", "build", "main.go"]