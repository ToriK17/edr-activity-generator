#!/bin/sh
set -e

# Detect distro for conditional logic on Fedora container
OS_ID=$(grep '^ID=' /etc/os-release | cut -d= -f2)

skip_uptime_hostname=false
if [ "$OS_ID" = "fedora" ]; then
  skip_uptime_hostname=true
fi

echo "=== Clean old logs ==="
./edr-activity-generator clean || true

echo "=== Test fallback process (sleep 1) ==="
./edr-activity-generator simulate process --count 1

echo "=== Test process: echo hello world ==="
./edr-activity-generator simulate process echo hello world

echo "=== Test process: whoami ==="
./edr-activity-generator simulate process whoami

echo "=== Test process: date ==="
./edr-activity-generator simulate process date

if [ "$skip_uptime_hostname" = false ]; then
  echo "=== Test process: uptime ==="
  ./edr-activity-generator simulate process uptime

  echo "=== Test process: hostname ==="
  ./edr-activity-generator simulate process hostname
else
  echo "=== Skipping uptime and hostname (not available in minimal Fedora image)"
fi

echo "=== Test process: shell-spawned ==="
./edr-activity-generator simulate process -- /bin/sh -c "echo from shell"

echo "=== Test simulate files (streaming 5s) ==="
./edr-activity-generator simulate files --stream 5s --delay 1s

echo "=== Test simulate process with YAML ==="
./edr-activity-generator simulate process --count 1 --format yaml echo YAMLtest

echo "=== Test simulate process with CSV ==="
./edr-activity-generator simulate process --count 1 --format csv echo CSVtest

echo "=== Test custom output path ==="
mkdir -p logs/testdir
./edr-activity-generator run --output logs/testdir/my_custom_logs.json

echo "=== Test simulate process with streaming ==="
./edr-activity-generator simulate process --stream 3s

echo "=== Run full simulation ==="
./edr-activity-generator run

echo "=== Test clean command ==="
./edr-activity-generator clean
ls logs || echo "Clean verified"

echo "=== All tests completed successfully. ==="
