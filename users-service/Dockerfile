FROM golang:alpine3.16 AS build2
LABEL stage=build2
WORKDIR /app
COPY . ./

RUN apk add build-base
RUN go build cmd/main.go

FROM alpine:3.16 AS runner
WORKDIR /app
COPY --from=build2 /app/main /app/main
COPY config.yml /app/config.yml
COPY /cert /app/cert
COPY env.example /app/env.example

CMD ["/app/main"]