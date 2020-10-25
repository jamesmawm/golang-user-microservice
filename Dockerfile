FROM golang:1.15-alpine

RUN apk add --update --no-cache gcc libc-dev

WORKDIR /go/src/app

COPY . .

EXPOSE 80

RUN go mod download && \
    go mod vendor && \
    go build -v && \
    rm -rf vendor && \
    mv golang-user-microservice server 

CMD [ "./server" ]
