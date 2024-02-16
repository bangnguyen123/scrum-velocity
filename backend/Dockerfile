FROM golang:1.21

WORKDIR /app

# Install Air for live-reloading
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
