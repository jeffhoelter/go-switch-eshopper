# Build environment
# -----------------
FROM golang:1.13-alpine as build-env
WORKDIR /go-switch-eshopper
RUN apk update && apk add --no-cache gcc musl-dev git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /go-switch-eshopper/bin/app /go-switch-eshopper/
COPY --from=build-env /go-switch-eshopper/bin/migrate /go-switch-eshopper/
COPY --from=build-env /go-switch-eshopper/migrations /go-switch-eshopper/migrations

EXPOSE 8080

CMD ["/go-switch-eshopper/app"]