FROM golang:1.21.3-alpine as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod ./

RUN go mod tidy && go mod download

COPY *.go ./

RUN go build -o /godocker

FROM alpine:latest

COPY --from=builder /godocker /godocker

EXPOSE 5000

CMD ["/godocker"]