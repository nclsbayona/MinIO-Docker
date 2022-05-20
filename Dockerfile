FROM golang:buster
WORKDIR /go/src/
RUN go mod init github.com/nclsbayona/minio_docker
COPY main.go .
RUN go mod tidy
RUN go mod vendor
ENTRYPOINT ["bash"]