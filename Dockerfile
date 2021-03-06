FROM golang:1.18 as builder
WORKDIR /usr/src/foo

COPY . .
RUN CGO_ENABLED=0 go build -o /foo ./foo

FROM gcr.io/distroless/base:debug
COPY --from=builder /foo /
ENTRYPOINT ["/foo"]
