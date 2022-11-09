FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /main

EXPOSE 5000

CMD [ "/main"]