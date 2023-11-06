FROM golang:1.21.3-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod ./

RUN go mod tidy && go mod download

COPY *.go ./

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker"]