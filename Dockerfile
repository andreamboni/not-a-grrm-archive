FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /not-a-grrm-archive

EXPOSE 8085

CMD ["/not-a-grrm-archive"]