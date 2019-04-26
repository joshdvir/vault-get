FROM golang:1.11

RUN mkdir -p /go/src/vault-get
WORKDIR /go/src/vault-get

EXPOSE 9292 2345
RUN curl https://glide.sh/get | sh
RUN go get github.com/derekparker/delve/cmd/dlv

ADD . /go/src/vault-get
RUN make deps
# RUN make build_inside_docker
CMD ["./vault-get"]