FROM golang:1.24-alpine

WORKDIR /app

# Install required tools for git and build
RUN apk add --no-cache git

COPY . .

# RUN go build -o edr-activity-generator .
RUN chmod +x docker/full-test.sh && \
    CGO_ENABLED=0 go build -o edr-activity-generator .

CMD ["./docker/full-test.sh"]