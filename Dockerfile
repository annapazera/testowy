FROM golang:1.11.10-alpine3.9 
RUN mkdir -p /src/testowy/
WORKDIR $GOPATH/src/testowy/
COPY . .
RUN apk update && apk add bash
RUN go build -o main .
CMD ["/testowy/main"]