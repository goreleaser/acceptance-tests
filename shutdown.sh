#!/bin/bash

set -e

WORK_DIR=$(pwd)

docker-compose -f docker-compose-gitea.yml down
echo "removing ssh key"
ssh-add -d ssh/id_rsa_goreleaser 
echo "removing directories"
rm -rf "$WORK_DIR"/gitea
rm -rf "$WORK_DIR"/ssh && mkdir -p "$WORK_DIR"/ssh