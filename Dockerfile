FROM golang:1.16.0 as builder

RUN apt-get update && apt-get install -y \
    bash-completion \
    --no-install-recommends && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/app
COPY . /go/src/app

RUN make setup-tools
RUN go build -o binary main.go

# COPY /go/src/app/binary /go/src/app/binary

# RUN go env -w GO111MODULE=auto
# RUN go build /app/main.go

# CMD ["go run main.go"]
CMD ["/go/src/app/binary"]