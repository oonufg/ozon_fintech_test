FROM golang:alpine AS build
WORKDIR /application
COPY /. .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0  go build -o main ./cmd/app


FROM alpine:latest
COPY --from=build /application/main .
COPY --from=build /application/.env .env
ENTRYPOINT ["./main"]

