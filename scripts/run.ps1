param (
    $command
)

if (-not $command)  {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

$env:XKAPUSTAJ_API_ENVIRONMENT="Development"
$env:XKAPUSTAJ_API_PORT="8086"

switch ($command) {
    "start" {
        go run ${ProjectRoot}/cmd/xkapustaj-api-service
    }
    "openapi" {
        docker run --rm -ti -v ${ProjectRoot}:/local openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
    }
    default {
        throw "Unknown command: $command"
    }
}