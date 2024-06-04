# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /uploadService

COPY . .

RUN go mod download

RUN go build -o bin .

EXPOSE 8180

ENTRYPOINT ["/uploadService/bin"]