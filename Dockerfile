FROM golang:1.13.7-alpine as builder

ADD . /app
RUN cd /app/cmd/server && go build -mod=vendor

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/server/server /app

CMD [ "/app/server" ]