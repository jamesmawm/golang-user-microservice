FROM golang:1.15-alpine

WORKDIR /go/src/app

COPY . .

EXPOSE 80

RUN go mod download && \
    go mod vendor && \
    go build -v && \
    rm -rf vendor && \
    mv golang-user-microservice server 

CMD [ "./server" ]
