FROM golang:1.10 AS build
RUN mkdir -p /app
COPY ./ /app
WORKDIR /app
RUN go get google.golang.org/grpc && go get github.com/garyburd/redigo/redis
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine
WORKDIR /app
COPY --from=build /app/ .
EXPOSE 9090
ENTRYPOINT [ "./server" ]