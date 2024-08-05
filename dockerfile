FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o main .

ENV PORT=2424

EXPOSE 2424

CMD ["/main"]

# FROM golang:1.22 as build
# WORKDIR /app
# COPY . .
# RUN go build -o /server .

# FROM scratch
# COPY --from=build /server /server
# EXPOSE 3000
# CMD ["/server"]