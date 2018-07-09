# build stage
FROM golang:1.10 AS build-env
RUN mkdir -p /go/src/github.com/evankanderson/efgh-samples/binary-input
WORKDIR /go/src/github.com/evankanderson/efgh-samples/binary-input
ADD . .
RUN go get -v ./...
# RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w' -o receiver

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/evankanderson/efgh-samples/binary-input/receiver /receiver
ENTRYPOINT ["/receiver"]
ENV PORT=8080
# Actually $PORT
EXPOSE 8080
