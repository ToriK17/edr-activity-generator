FROM golang:1.24-bookworm

WORKDIR /app

# Copy source code
COPY . .

# Build the binary
# RUN go build -o edr-activity-generator .
RUN CGO_ENABLED=0 go build -o edr-activity-generator .


# Default run command
CMD ["./edr-activity-generator", "run"]
