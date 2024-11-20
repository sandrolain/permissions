#!/bin/bash

cmd=$()

hyperfine "buf curl --http2-prior-knowledge --protocol grpc --data '{\"user\": \"foo\", \"scope\": \"bar\"}' http://localhost:9090/permissionsGrpc.PermissionsService/UserAllowed"