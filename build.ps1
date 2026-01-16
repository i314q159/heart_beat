$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o heart_beat_amd64 main.go
$env:GOOS="linux"; $env:GOARCH="arm64"; go build -o heart_beat_arm64 main.go
