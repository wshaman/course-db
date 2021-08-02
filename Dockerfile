FROM golang:1.16-alpine as builder
COPY ./src /src
WORKDIR /src
RUN mkdir -p /src/bin/
RUN ls
RUN go build ./cmd/migrate -o /src/bin/migrate
RUN go build ./cmd/api -o /src/bin/api

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /src/bin/migrate /app/migrate
COPY --from=builder /src/bin/api /app/api
ENTRYPOINT ["bash"]




