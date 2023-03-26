FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

ENV VERBOSE=1

COPY ./ ./

# RUN go build -o build ./src/main.go

EXPOSE 50051

ENTRYPOINT [ "go", "run", "src/main.go" ]