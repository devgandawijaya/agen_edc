# build stage
FROM golang:1.23-alpine AS build
WORKDIR /app
# install git only if `go get` needs it (not needed here), keep minimal
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/server ./cmd/server

# final stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=build /app/bin/server .
ENV PORT=2021
EXPOSE 2021
CMD ["./server"]
