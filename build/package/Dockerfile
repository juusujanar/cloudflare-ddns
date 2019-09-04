FROM golang:1.13-alpine as builder
WORKDIR /src/github.com/juusujanar/cloudflare-ddns

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates git

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
ENV CGO_ENABLED=0 GOOS=linux

COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build -o /app ./cmd/cloudflare-ddns

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