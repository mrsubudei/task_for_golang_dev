FROM golang:alpine3.16 AS build1
LABEL stage=build1
WORKDIR /app
COPY . ./

RUN apk add build-base
RUN go build cmd/grpc/main.go

FROM alpine:3.16 AS runner
WORKDIR /app
COPY --from=build1 /app/main /app/main
COPY config.yml /app/config.yml
COPY /cert /app/cert

CMD ["/app/main"]