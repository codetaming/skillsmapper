FROM golang:1.15.5-alpine3.12 AS BUILD
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o skillsmapperd ./cmd/skillsmapperd

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
RUN mkdir /app
ENV SERVER_PORT=8080
COPY --from=BUILD /app/skillsmapperd .
CMD ["./skillsmapperd"]
EXPOSE 8080