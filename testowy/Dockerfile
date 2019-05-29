FROM golang:1.11.10-alpine3.9
RUN mkdir -p /src/testowy/
ADD . /src/testowy/
WORKDIR /src/testowy/
RUN apk update && apk add bash
RUN go build -o main .
CMD ["/src/testowy/main"]