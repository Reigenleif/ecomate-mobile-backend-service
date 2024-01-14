FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
COPY / ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

RUN go mod download 
# RUN go mod tidy
RUN go get github.com/Reigenleif/ecomate-mobile-backend-service
# RUN go build 

CMD ["go", "run", "main.go"]