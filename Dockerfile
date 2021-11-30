FROM golang:1.17 as builder
WORKDIR /go/src/gdhuo
COPY . .

ENV GO111MODULE=on

RUN go mod download
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w" -o app main.go
RUN cp app /go/bin/app

FROM alpine:latest as builder2
RUN apk add --no-cache upx

COPY --from=builder /go/bin/app /go/bin/app
WORKDIR /go/bin
RUN upx app
RUN apk del --no-cache upx

FROM scratch
# Copy our static executable.
COPY --from=builder2 /go/bin/app /
# Copy openapi spec
COPY ./openapi.spec.yaml /doc/openapi.spec.yaml
ENV SPEC_PATH=/doc/openapi.spec.yaml
# Run the hello binary.
ENV ADDRESSES=:8080
EXPOSE 8080
ENTRYPOINT ["/app"]
