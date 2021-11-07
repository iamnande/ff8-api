# container: source and metadata
FROM scratch

# container: add compiled binary to image
COPY build/api /api

# container: run
ENTRYPOINT ["/api"]