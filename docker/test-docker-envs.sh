#!/bin/bash

set -e

for distro in debian alpine fedora; do
  echo "=== Building and running in $distro ==="
  docker build -t edr-test-$distro -f docker/$distro/Dockerfile .
  docker run --rm edr-test-$distro
done
