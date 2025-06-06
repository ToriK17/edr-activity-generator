# EDR Activity Generator — README

[For a brief one-page project overview](/project-overview.md)

## Purpose

This tool is a cross-platform command-line tool that simulates realistic endpoint activity (processes, files, and network traffic) and produces structured telemetry logs. It's designed to help validate the accuracy of telemetry collected by Endpoint Detection and Response (EDR) agents, especially after updates—to catch regressions before they impact production.

## Why Go?

Although my background is primarily in Ruby, I chose Go for its excellent cross-platform capabilities, native system-level libraries, and ability to compile static binaries. This project provided a valuable opportunity to deepen my Go experience while solving a practical, low-level problem in a performance-conscious language.

## Key Features

- **Cross-platform support:** Tested on macOS and Linux.
- **Flexible output formats:** Supports JSON (default), YAML, and CSV.
- **Activity types simulated:**
  - Process creation (`sleep 1`)
  - File creation, modification, and deletion
  - Network activity
    - HTTP/1.1 via raw TCP socket (example.com:80)
    - HTTP/2 via HTTPS request (https://nghttp2.org)
- **Structured telemetry output:** One log entry per action, machine-ingestible and human-readable:

## Usage

#### Build

`go build -o edr-activity-generator .`

You can invoke the tool via the CLI using the following subcommands:

#### Run All Simulations

`./edr-activity-generator run`
Generates logs in logs/activity_log.json by default.

#### Run Individual Simulations
`./edr-activity-generator simulate process --count 5 --format yaml`
`./edr-activity-generator simulate files --stream 10s --delay 1s`
`./edr-activity-generator simulate network --format csv`

#### Clean the Log File
`./edr-activity-generator clean`
Removes logs/activity_log.json if it exists.

#### Generate Activity and customize the output path & directories

`./edr-activity-generator run --output logs/custom_activity_log.json`
` ./edr-activity-generator run --output new_place/my_custom_logs.json`

## Log Format

Each entry includes relevant telemetry fields depending on the activity type.

#### Common Fields

- timestamp (RFC3339)
- username
- process_name
- command_line
- process_id

#### File Activity

- file_path
- action: one of create, modify, delete

#### Network Activity

- source_address
- destination_address
- protocol
- bytes_sent

## Testing

Basic unit tests are included for the log writing and CSV serialization logic (activity/log_writer_test.go). The code is modular and testable, with room for expansion if deeper validation or mocks are needed in the future.

## Container Support

This tool has been tested in:

- Debian
- Alpine
- Fedora

Each has a corresponding Dockerfile under `docker/`. A helper script is available to validate builds across environments:
`./docker/test-docker-envs.sh`

## Potential Future Improvements

This project focused on delivering a clear, functioning MVP. That said, I made several intentional scope decisions with an eye toward future extensibility. Here are the planned enhancements if time was not a factor.

#### Platform Support & Portability

- **Windows Support:** Extend compatibility to Windows systems with proper API handling (e.g., via syscall or cross-platform abstractions).
- **Platform Detection:** Automatically adapt behavior using runtime.GOOS and annotate log entries with platform metadata.

#### Enhanced Realism & Load Simulation

- **Simulated Bursts of Activity:** Allow bursts of activity to better reflect real-world usage spikes (`--high-cpu`, `--cstream`).
- **Concurrent Simulations:** Leverage Go’s goroutines to simulate multiple activity types in parallel, improving realism and testing concurrency.
- **Stress Testing Modes:** Optional flags to simulate elevated system load for process creation or file operations, clearly labeled and designed for isolated, sandboxed environments only. These would help test how EDR agents behave under pressure, without introducing real system risk.

#### Logging Architecture

- **Centralized, Non-Blocking Logging:** Use channels to implement a dedicated logging goroutine for thread-safe, performant output handling.
- **Expanded Telemetry Fields:** Include source and destination ports in network activity logs; include GOOS in all entries for platform traceability.
- **Improved Built-in Formatting:** Add human-readable formatting options (e.g., pretty-print JSON) to avoid reliance on external tools like jq.

#### Testability & CI Integration

- **Comprehensive Test Coverage:** Expand testins to include tests for each simulation type.
- **Optional CI Pipeline Integration:** Design the tool to run in a CI job and validate agent output consistency post-deploy.

#### Telemetry Validation & Regression Detection

- **Agent Log Comparison Tooling:** Build optional validators to compare generated activity with agent-emitted telemetry (out of scope for this demo, but aligned with long-term goals).

- **Failure Alerts:** Optional flags to mark missing or malformed telemetry if paired with EDR logs, enabling early detection of regression bugs.

## Additional Feature Considerations

If this tool is used to validate agent behavior post-update, these improvements may also help:

- **Custom Activity Profiles:** Load activity instructions from a config file (YAML/JSON) to simulate varied test scenarios across runs.

- **Timestamp Jitter/Delay Injection:** Simulate real-world delays in logging or activity to test how EDR agents handle imperfect timing.

- **Replay Mode:** Save & re-run exact sequences of generated activity to help reproduce regressions from real test cases.

- **Telemetry Schema Validator:** Lightweight JSON schema checks to flag missing or incorrect fields in generated logs, useful during pipeline validation.

## Scope Prioritization

Several advanced features (Windows support, EDR validation logic) were deferred to focus on delivering a clear and reviewable MVP. Each planned feature reflects either specific feedback or practical considerations for future-proofing the tool.
