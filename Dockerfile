FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/mhurtado/go-api-example
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/mhb-go-api-example ./cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/mhb-go-api-example /go/bin/mhb-go-api-example
ENTRYPOINT ["/go/bin/mhb-go-api-example"]
