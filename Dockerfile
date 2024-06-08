FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go get -u github.com/99designs/gqlgen@latest
RUN go get -u github.com/99designs/gqlgen/codegen/config@latest
RUN go get -u github.com/99designs/gqlgen/internal/imports@latest
RUN go get -u github.com/urfave/cli/v2@latest

ENV GO111MODULE=on

COPY . .

RUN go mod tidy

RUN go run github.com/99designs/gqlgen generate

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
