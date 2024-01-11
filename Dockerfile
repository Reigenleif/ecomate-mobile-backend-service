FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin

ENTRYPOINT [ /app/bin ]