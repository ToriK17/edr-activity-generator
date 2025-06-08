# Build stage 1 with all necessary GO tooling in a linux env
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN chmod +x docker/full-test.sh && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o edr-activity-generator

# CGO_ENABLED=0 means that the binary avoids depending on system C libs, it becomes self-contained, minimizing distro-specific breakage.

# Stage 2 Just copies in the precompiled binary and runs it in fedora 
FROM fedora:38
WORKDIR /app

COPY --from=builder /app/edr-activity-generator .
COPY --from=builder /app/docker/full-test.sh ./docker/full-test.sh
RUN chmod +x ./docker/full-test.sh
CMD ["./docker/full-test.sh"]

# Built on Debian but ran in Fedora so Fedora’s libc, kernel, and runtime environment do support the binary.