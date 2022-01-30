FROM golang:alpine
WORKDIR /go/src/app
COPY ./server .
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "server.go", "-env", "dev"]

