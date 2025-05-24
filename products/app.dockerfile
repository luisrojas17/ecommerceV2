FROM golang:1.24.3-alpine3.21 as build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/luisrojas17/ecommerceV2
COPY go.mod go.sum ./
COPY vendor vendor
COPY products products
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./products/cmd

FROM alpine:3.21
WORKDIR /usr/bin
COPY --from=build /go/bin .
CMD ["app"]