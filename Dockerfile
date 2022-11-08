FROM golang:1.18.7-alpine as builder


COPY . /go/src/stuart.com/gocmd
RUN cd /go/src/stuart.com/gocmd && CGO_ENABLED=0 go build -o /go/bin/gocmd

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/gocmd  /gocmd

EXPOSE 8982

# Run the hello binary.
ENTRYPOINT ["/gocmd"]
