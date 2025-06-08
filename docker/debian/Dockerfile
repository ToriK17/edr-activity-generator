FROM golang:1.24-bookworm

WORKDIR /app

# Copy source code
COPY . .

# Build the binary
# RUN go build -o edr-activity-generator .
RUN chmod +x docker/full-test.sh && \
    CGO_ENABLED=0 go build -o edr-activity-generator .

CMD ["./docker/full-test.sh"]
