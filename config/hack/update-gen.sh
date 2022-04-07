#!/bin/bash

BASE="github.com/slabs-forge/klabs-api"
HEADER=hack/header.go.txt
s
deepcopy-gen --input-dirs ${BASE}/apis/slabs.forge/v1 -O zz_generated.deepcopy --trim-path-prefix ${BASE} -h ${HEADER}
