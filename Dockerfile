

FROM golang:1.22.3
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker_to_do_cli
EXPOSE 8080

# Run
CMD ["/docker_to_do_cli"]



