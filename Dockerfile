FROM golang:1.21-alpine AS build

# Set the GOPATH environment variable
ENV GOPATH /go

# Create the workspace
RUN mkdir -p $GOPATH/src $GOPATH/bin

# Copy project files
COPY ./*.go $GOPATH/src
COPY ./go.mod $GOPATH/src
COPY ./namesgenerator $GOPATH/src/namesgenerator
COPY ./public $GOPATH/src/public
COPY ./static $GOPATH/src/static

RUN ls -la $GOPATH/src/public/views

# Set working directory
WORKDIR $GOPATH/src

RUN go mod tidy
#RUN go install
#RUN go get rolling-code-generator

# Build the application
RUN go build -v -o $GOPATH/bin/app

FROM scratch
WORKDIR /usr/local/bin
COPY --from=build /go/bin/app /usr/local/bin/app
COPY --from=build /go/src/public /usr/local/bin/public
COPY --from=build /go/src/static /usr/local/bin/static

ENTRYPOINT ["/usr/local/bin/app"]
