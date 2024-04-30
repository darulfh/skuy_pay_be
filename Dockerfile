FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o main .

ENV PORT=2424

EXPOSE 2424

CMD ["./main"]
