#!/usr/local/bin/pwsh
Param($target=$null)

$commands = @("./cmd/deployctrl", "./cmd/mxdeployservice")

if (-not $target) {
    $goenv = go env GOOS GOARCH
    $targets = @("$($goenv[0])/$($goenv[1])")
} elseif ($target -eq "all") {
    $targets = @("darwin/amd64", "darwin/arm64", "windows/amd64")
} else {
    $targets = @($target)   
}

if ((Test-Path ./dist)) { Remove-Item -r -fo ./dist }

foreach ($target in $targets) {
    $goenv = $target.Split("/")
    $env:GOOS = $goenv[0]
    $env:GOARCH = $goenv[1]
    
    $dest = "./dist/$($env:GOOS)/$($env:GOARCH)"
    New-Item -ItemType Directory $dest | Out-Null

    foreach ($command in $commands) {
        go build -o $dest $command
    }
}
