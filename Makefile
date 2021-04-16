ROOTDIR := $(shell pwd)
OUTPUT_DIR=$(ROOTDIR)/_output

GCFLAGS := all="-N -l"

APISERVER_SOURCES=$(shell find $(ROOTDIR) -name "*.go" | grep -v test | grep -v generated.go)

$(OUTPUT_DIR)/apiserver-linux-amd64: $(APISERVER_SOURCES)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -gcflags=$(GCFLAGS) -o $@ $(ROOTDIR)/apiserver/cmd/main.go

.PHONY: apiserver-docker
apiserver-docker: $(OUTPUT_DIR)/apiserver-linux-amd64
	docker build -t ashwang168/capture-the-flag-apiserver:v1 $(OUTPUT_DIR) -f $(ROOTDIR)/apiserver/cmd/Dockerfile;

# Run the apiserver on localhost:1234
.PHONY: run-apiserver
run-apiserver:
	go run apiserver/cmd/main.go

.PHONY: release
release: apiserver-docker
	docker push ashwang168/capture-the-flag-apiserver:v1

# Clean
.PHONY: clean
clean:
	rm -rf _output
	rm -rf vendor_any
	rm -rf pkg/api