# capture-the-flag-apiserver

## Scenario
A company that collects lots of data stores it in Kubernetes, but since they want to show their users some information,
they provide an apiserver (that's really insecure). Later, they publish the protos that they used to build the
apiserver online.

See the protos at `api/apiserver.proto`

## Task
Build a grpc client that can read any data from the apiserver, using the CaptureTheFlag rpc call.