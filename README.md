# capture-the-flag-apiserver

## Scenario
A company that collects lots of data stores it in Kubernetes, but since they want to show their users some information,
they provide an apiserver (that's really insecure). Later, they publish the protos that they used to build the
apiserver online.

See the protos at `api/apiserver.proto`

## Prerequisites
Have the following tools available:
* go 1.16
* grpcurl

Download this repo:
```shell script
git clone https://github.com/ashleywang1/capture-the-flag-apiserver.git
cd capture-the-flag-apiserver
```

## Setup/Tutorial
Run:
```shell script
make run-apiserver
```

In a separate terminal, run:
```
grpcurl -plaintext localhost:1234 list
```

You should see the following output:
```
{
  "flag": {
    "flag": "[the flag will be here]"
  }
}
```

## Task
Build a grpc client that can read any data from the apiserver, using the CaptureTheFlag rpc call.

Point to the official apiserver url given to you, and capture the flag!
