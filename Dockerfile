FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

SHELL ["/bin/bash", "-c", "source .env"]

CMD [ "make", "run" ]