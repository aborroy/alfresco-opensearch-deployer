#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

# Cleanup resources
function cleanup {
  set +e
  docker compose down
  cd ..
  rm -rf output
  set -e
}

function waitAlfrescoReady {
  echo "Starting Alfresco ..."
  bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' http://localhost:8080/alfresco/s/api/server)" != "200" ]]; do sleep 5; done'
  echo "Alfresco started successfully!"
}

cd ..

# Alfresco Community
go run main.go create -i=false -v 2.13.0

cd output

docker compose up -d

waitAlfrescoReady

cleanup