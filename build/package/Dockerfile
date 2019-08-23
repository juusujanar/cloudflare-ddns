FROM golang:1.12-alpine as builder
WORKDIR $GOPATH/src/github.com/juusujanar/cloudflare-ddns

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates git

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
# * GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder.
ENV CGO_ENABLED=0 GOOS=linux GOFLAGS=-mod=vendor

COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN dep ensure --vendor-only && \
    go build \
    -installsuffix 'static' \
    -o /app ./cmd/cloudflare-ddns


# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the second stage.
COPY --from=builder /app /app

# Perform any further action as an unprivileged user.
USER 65534:65534

# Run the compiled binary.
ENTRYPOINT ["/app"]