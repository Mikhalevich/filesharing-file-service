FROM golang:latest as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/file_service

FROM scratch
COPY --from=builder /go/bin/file_service /go/bin/file_service

EXPOSE 50051

WORKDIR /go/bin
ENTRYPOINT ["./file_service"]
