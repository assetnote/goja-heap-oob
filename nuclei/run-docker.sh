#!/usr/bin/env bash

HERE=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)
docker run -it -v "${HERE}:/app" -w /app ubuntu:26.10 /bin/bash