########################################################
# Using multi-stage build to reduce deployment image size
########################################################
########################################################
# STEP 1 build executable binary
########################################################
# Because official golang image size is bigger than golang:alpine
FROM golang:1.18.1-alpine AS builder

WORKDIR $GOPATH/src/deploy-aws/
COPY . .

# Fetch depnedancies. Must cd to the main package before fetching.
# -d stop after downloading packages
# -v verbose progress and debug output
RUN cd cmd && go get -d -v

# Remove debug information, compile only for linux target, disable cross compilation
# go version < 1.10
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/deploy-aws
# go version >= 1.10
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/deploy-aws cmd/main.go

########################################################
# STEP 2 build a small image
########################################################
FROM gcr.io/distroless/static:latest-arm64

# Copy static executable
COPY --from=builder --chown=nonroot:nonroot /go/bin/deploy-aws /go/bin/deploy-aws


# Execute as nonroot
USER nonroot

# Listening on port 8088 for client, 8089 for operator
EXPOSE 8088
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002

# Run the application
ENTRYPOINT ["/go/bin/deploy-aws"]