FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /redikru-server ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /redikru-server /redikru-server
EXPOSE 8081
CMD ["/redikru-server"]
