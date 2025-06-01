# EDR Activity Generator — Project Overview

## Purpose

This tool generates real endpoint activity across supported platforms (macOS and Linux in this implementation) to help detect regressions in telemetry emitted by Endpoint Detection and Response (EDR) agents after updates. It provides structured logging of triggered actions to support correlation with what the EDR agent records.

## Why Go?

Although I’m more experienced in Ruby, I chose to build this tool in Go because it aligns with Red Canary’s internal stack and offers robust cross-platform support with strong system-level access. This project has been a great opportunity to sharpen my Go skills and demonstrate my ability to learn and adapt quickly in service of a real-world use case.

## Key Features

- **Cross-platform support:** macOS and Linux currently tested.
- **Modular CLI built with Cobra:** Includes subcommands for `run` and `clean`, with room for expansion.
- **Activity simulation includes:**
  - Process creation (`sleep 1` as a placeholder)
  - File creation, modification, and deletion
- **Structured telemetry log:** Each activity is written as a separate JSON object on its own line:
  - `timestamp`
  - `username`
  - `process_name`
  - `command_line`
  - `process_id`
  - `file_path` and `action` (present for file-related activity only)

## How It Works

### Locally Build the CLI

`go build -o edr-activity-generator .`

You can invoke the tool via the CLI using the following subcommands:

### Generate Activity

`./edr-activity-generator run`
The above defaults to logs/activity_log.json

### Generate Activity and customize the output path

`./edr-activity-generator run --output logs/<custom_activity_log>.json`

### Clean Existing Log File

`./edr-activity-generator clean`
