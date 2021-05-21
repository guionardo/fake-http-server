FROM golang:1.16
WORKDIR /go/src/github.com/guionardo/fake-http-server/
RUN go get -d -v github.com/gin-gonic/gin
ARG GIN_MODE
COPY fake-http-server.go .
COPY go.* ./
RUN CGO_ENABLED=0 GOOS=linux go build -o fake-http-server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/guionardo/fake-http-server/fake-http-server .
CMD ["./fake-http-server"]
