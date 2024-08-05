FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o main .

ENV PORT=8080

EXPOSE 8080

CMD ["/main"]