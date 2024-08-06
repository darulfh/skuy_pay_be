FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o server .

# FROM scratch
# COPY --from=build /server /server

ENV PORT=2424

EXPOSE 2424

CMD ["./server"]

# FROM golang:1.22 as build
# WORKDIR /app
# COPY . .
# RUN go build -o /server .

# FROM scratch
# COPY --from=build /server /server
# EXPOSE 2424
# CMD ["/server"]