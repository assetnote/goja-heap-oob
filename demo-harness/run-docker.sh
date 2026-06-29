#!/usr/bin/env bash

HERE=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)
docker run -it -v "${HERE}:/app" -w /app golang:1.25 /bin/bash