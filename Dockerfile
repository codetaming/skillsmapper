FROM golang:1.13-stretch AS BUILD
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o skillsmapper .

FROM alpine:3.8
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir /app
ENV SERVER_PORT=8080
COPY --from=BUILD /app/skillsmapper .
CMD ["./skillsmapper"]
EXPOSE 8080