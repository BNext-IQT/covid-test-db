FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
COPY ./models /models
WORKDIR /build

# Copy and download dependency using go mod
COPY api/go.mod .
COPY api/go.sum .
RUN go mod download

# Copy the code into the container
COPY api .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 5000

# Command to run when starting the container
CMD ["/dist/main"]