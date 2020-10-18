FROM golang:1.15-alpine

WORKDIR /go/src/app

COPY . .

EXPOSE 80

RUN go mod init v1 && \
    go mod download && \
    go mod vendor && \
    go build -v && \
    rm -rf vendor && \
    mv v1 server 

CMD [ "./server" ]