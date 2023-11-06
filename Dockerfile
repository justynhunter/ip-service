FROM golang:1.21-alpine as base

### BUILD ###
FROM base as builder
WORKDIR /usr/build

COPY . .
RUN go build

### RUN ###
FROM base as runner

ARG PORT

WORKDIR /usr/app
COPY --from=builder /usr/build/ip-service ./

EXPOSE ${PORT}

CMD ["ip-service"]
