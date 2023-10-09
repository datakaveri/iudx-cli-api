FROM golang:1.21.2-alpine AS build

RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/main /usr/bin/
ENTRYPOINT [ "main" ]