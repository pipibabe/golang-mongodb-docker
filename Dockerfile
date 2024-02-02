FROM golang:latest

WORKDIR /Users/pipi/projects/golang/golang-mongodb-docker

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./golang_test .
RUN go build -v -o ./golang_test ./golang_test/build

CMD ["./golang_test/build/app"]