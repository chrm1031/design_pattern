FROM golang:1.18


ENV DIR=/usr/src/cmd
WORKDIR ${DIR}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
# COPY go.mod ./
# RUN go mod download && go mod verify
# RUN go mod init example.com/go-mod-test