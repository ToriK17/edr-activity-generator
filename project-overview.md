# EDR Activity Generator — Project Overview

This project implements a cross-platform command-line tool designed to simulate realistic endpoint activity and generate structured telemetry logs. The purpose is to help validate that updated EDR (Endpoint Detection and Response) agents continue to capture key behavioral data consistently and accurately to prevent potential regressions.

## What It Does

### The tool simulates the following endpoint activities:
- Process creation (sleep 1)
- File operations: creation, modification, and deletion
- Network activity: Basic HTTP/1.1 traffic via raw TCP socket & HTTP/2 traffic via HTTPS requests to nghttp2.org

Each action is logged in a machine-ingestible format (default: JSON), and added support for YAML and CSV output included via flags.

### The logs capture essential metadata:
- Timestamps, User and process information, File paths and actions, Network source/destination addresses and protocols

## CLI Design
The tool is built using Cobra for a modular, extensible CLI interface.

### Available Commands
***Run all activity types at once (default JSON):***
- `./edr-activity-generator run`

***Run individual types:***
- `./edr-activity-generator simulate process --count 5 --format yaml`

***Stream activities over time:***
- `./edr-activity-generator simulate network --stream 10s --format csv --delay 1s`

***Clean log output.***
- `./edr-activity-generator clean`

## Testing
The project includes a lightweight unit tests for CSV encoding, located in activity/log_writer_test.go. To keep scope focused on cross-platform simulation and telemetry structure, I opted not to include a larger test suite. However, individual components are modular and could be easily tested using standard Go test tooling.

### Container Support
Tested and runnable in: Debian, Alpine, and Fedora.

Each distro has a corresponding Dockerfile. A `test-docker-envs.sh` script is included to build and run the container in all three environments, verifying the binary’s portability.

### Platforms Supported
Tested on macOS and Linux. The project avoids CGO and uses cross-platform-safe system packages (`os/exec`, `os`, `net`, etc.) to produce statically linked binaries for portability.

## Future Considerations

If further developed, the tool could include:
- Windows support and platform-aware behavior
- Simulated concurrency to better reflect real-world endpoint load
- Optional telemetry validation tooling to compare agent output against known activity
- Configurable profiles to replay test scenarios or simulate high-volume activity

This project aimed to strike a balance between practicality, portability, and forward-thinking design, creating a reliable base for future test automation or integration into CI workflows.
