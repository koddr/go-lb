FROM golang:1.17-alpine AS builder

LABEL maintainer="Vic Shóstak <vic@shostak.dev> (https://shostak.dev/)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environmet variables needed for our image and build the load balancer server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o lb .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/lb", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/lb"]