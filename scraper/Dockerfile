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
COPY scraper/go.mod .
COPY scraper/go.sum .
RUN go mod download

# Copy the code into the container
COPY scraper .

# Build the application
RUN go build -o scraper .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist
COPY scraper/Database_Master.xlsx .

# Copy binary from build to main folder
RUN cp /build/scraper .

# Command to run when starting the container
CMD ["/dist/scraper"]