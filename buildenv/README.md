# Build Environment
We have a docker-based build environment which helps you to build the checkmake binary on a system that has no golang, etc. installed.

# Generate Build Environment
Run:
```
cd buildenv
docker build -t checkmake/buildenv:latest .
cd ..
```
to generate a build environment (docker image). This docker images can be used to build your checkmake binary. 

# Build Checkmake Binary
Run 
```
docker run --rm -v $(pwd):/data --workdir /data checkmake/buildenv:latest make
```
to generate the checkmake binary.
Output is similar to:
```
Checking the programs required for the build are installed...
install -d .d
echo "checkmake: $(go list -f '{{ join .Deps "\n" }}' cmd/checkmake/main.go | awk '/github/ { gsub(/^github.com\/[a-z]*\/[a-z]*\//, ""); printf $0"/*.go " }')" > .d/checkmake.d
go build -ldflags "-X 'main.version=0.1.0-22-g42f1561' -X 'main.buildTime=2020-02-23T16:02:51Z' -X 'main.builder= <>' -X 'main.goversion=go version go1.13.8 linux/amd64'" -o checkmake cmd/checkmake/main.go
sed "s/REPLACE_DATE/February 23, 2020/" man/man1/checkmake.1.md | pandoc -s -t man -o checkmake.1
```

# Test Checkmake Binary
Run checkmake binary:
```
./checkmake --version
```
to test the binary. Output should be similar to:
```
checkmake 0.1.0-22-g42f1561 built at 2020-02-23T16:02:51Z by  <> with go version go1.13.8 linux/amd64
```