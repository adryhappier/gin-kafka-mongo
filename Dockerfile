# Write A Dockerfile With Docker Multi-Stage Builds
# build stage
FROM golang:latest AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# working directory
WORKDIR /go/src/github.com/adryhappier/docker_imgs

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
# rebuilt built in libraries and disabled cgo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# Build the application
RUN go build -o main .

# final stage
FROM alpine:latest

# working directory
WORKDIR /go/src/github.com/adryhappier/docker_imgs

# Copy binary from build to main folder
# RUN cp /build/main .
# copy the binary file into working directory
COPY --from=builder /go/src/github.com/adryhappier/docker_imgs/main .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["./main"]


