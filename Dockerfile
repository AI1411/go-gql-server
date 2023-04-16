FROM golang:1.19.1

ARG USERNAME=AI1411
# hadolint ignore=DL3008

ENV TZ Asia/Tokyo

ENV GO111MODULE on
WORKDIR /go/src/
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]