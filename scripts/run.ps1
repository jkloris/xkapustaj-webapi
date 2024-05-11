param (
    $command
)

if (-not $command)  {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

$env:XKAPUSTAJ_API_ENVIRONMENT="Development"
$env:XKAPUSTAJ_API_PORT="8086"
$env:AMBULANCE_API_MONGODB_USERNAME="root"
$env:AMBULANCE_API_MONGODB_PASSWORD="neUhaDnes"

function mongo {
    docker compose --file ${ProjectRoot}/deployments/docker-compose/compose.yaml $args
}

switch ($command) {
    "start" {
        try {
            mongo up --detach
            go run ${ProjectRoot}/cmd/xkapustaj-api-service
        } finally {
            mongo down
        }
    }
    "mongo" {
        mongo up
    }
    "docker" {
        docker build -t jergusk/xkapustaj-wl-webapi:local-build -f ${ProjectRoot}/build/docker/Dockerfile .
    }
    "openapi" {
        docker run --rm -ti -v ${ProjectRoot}:/local openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
    }
    default {
        throw "Unknown command: $command"
    }
}