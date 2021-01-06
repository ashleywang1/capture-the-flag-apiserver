module github.com/ashleywang1/capture-the-flag-apiserver

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	github.com/mitchellh/hashstructure v1.0.0
	github.com/rotisserie/eris v0.5.0
	github.com/solo-io/go-utils v0.20.1
	github.com/solo-io/protoc-gen-ext v0.0.13
	github.com/solo-io/skv2 v0.15.2
	go.opencensus.io v0.22.5
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	k8s.io/apimachinery v0.18.6
)

replace (
	// github.com/Azure/go-autorest/autorest has different versions for the Go
	// modules than it does for releases on the repository. Note the correct
	// version when updating.
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.8.3 // indirect
	// kube 0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
)
