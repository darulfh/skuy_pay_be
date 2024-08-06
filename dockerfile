# FROM golang:1.22

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . /app

# RUN go build -o server .



# ENV PORT=2424

# EXPOSE 2424

# CMD ["./server"]

# FROM golang:1.22 as build
# WORKDIR /app
# COPY . .
# RUN go build -o /server .

# FROM scratch
# COPY --from=build /server /server
# EXPOSE 2424
# CMD ["/server"]


# Build stage
FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o coolify .

# Cache stage
FROM alpine:latest AS cache

WORKDIR /app

COPY --from=build /app/coolify .

# Final stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates curl

COPY --from=cache /app/coolify .

EXPOSE 2424

CMD ["./coolify"]