#!/bin/bash

BASE="github.com/slabs-forge/klabs-api"
APIS=${BASE}/apis/slabs.forge/v1
HEADER=hack/header.go.txt

deepcopy-gen --input-dirs ${APIS} -O zz_generated.deepcopy \
    --trim-path-prefix ${BASE} -h ${HEADER}

client-gen --clientset-name versioned --input-base "" --input ${APIS} --output-package ${BASE}/clientset \
    --trim-path-prefix ${BASE} -h ${HEADER}

lister-gen --input-dirs ${APIS} --output-package  ${BASE}/listers \
     --trim-path-prefix ${BASE} -h ${HEADER}

informer-gen --versioned-clientset-package ${BASE}/clientset/versioned --input-dirs ${APIS}  \
    --output-package ${BASE}/informers --listers-package ${BASE}/listers \
    --trim-path-prefix ${BASE} -h ${HEADER}
