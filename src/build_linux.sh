sudo apt-get install -y libgl1-mesa-dev xorg-dev gcc
CGO_ENABLED=1 go build -o LogCHEK-GUI main.go
CGO_ENABLED=1 go build -o LogCHEK cli.go
