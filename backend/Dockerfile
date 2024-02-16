FROM golang:1.21

WORKDIR /app

# Install Air for live-reloading
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY schema.prisma ./
RUN go run github.com/steebchen/prisma-client-go generate

CMD ["air", "-c", ".air.toml"]
