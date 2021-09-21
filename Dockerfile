#FROM golang:1.17
FROM golang:alpine3.14
COPY . .
RUN go build -o server main.go
CMD ["./server"]