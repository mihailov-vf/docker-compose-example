FROM golang:alpine as build

WORKDIR /go/src/app

ENV GO111MODULE=auto

COPY . .

RUN go build -v -o hello

FROM alpine

WORKDIR /app

COPY --from=build /go/src/app/ /app

RUN chmod -R 755 /app

EXPOSE 8080

ENTRYPOINT [ "/app/hello" ]