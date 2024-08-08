FROM golang:1.21-alpine AS build

# Set the GOPATH environment variable
ENV GOPATH=/go

# Create the workspace
RUN mkdir -p $GOPATH/src $GOPATH/bin

# Copy project files
COPY ./*.go $GOPATH/src
COPY ./go.mod $GOPATH/src
COPY ./cmd $GOPATH/src/cmd
COPY ./internal $GOPATH/src/internal
COPY ./web $GOPATH/src/web

# Set working directory
WORKDIR $GOPATH/src

# Tidy up and get dependencies
RUN go mod tidy

# Build the application
RUN go build -v -o $GOPATH/bin/app

# Use multi-stage build
FROM scratch
WORKDIR /usr/local/bin
COPY --from=build /go/bin/app /usr/local/bin/app
COPY --from=build /go/src/web /usr/local/bin/web

# Run the app
ENTRYPOINT ["/usr/local/bin/app"]
