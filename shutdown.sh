#!/bin/bash

set -e

WORK_DIR=$(pwd)

docker-compose -f docker-compose-gitea.yml down
echo "removing dirs"
rm -rf "$WORK_DIR"/gitea
rm -rf "$WORK_DIR"/ssh && mkdir -p "$WORK_DIR"/ssh