#!/bin/bash

set -e

for distro in debian alpine fedora; do
  echo "=== Building and running in $distro ==="
  docker build -t edr-test-$distro -f docker/$distro/Dockerfile .
  docker run --rm edr-test-$distro
done
# The Fedora container requires --security-opt seccomp=unconfined due to dnf’s syscall usage.
# This works for locally simulating activity, never for production