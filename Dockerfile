FROM golang:1.11.10-alpine3.9
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN apk update && apk add bash
RUN go build -o main .
CMD ["/src/main"]