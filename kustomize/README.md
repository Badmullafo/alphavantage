# kustomize

<!-- ABOUT THE PROJECT -->
## About The Project

This readme provides instructions on how to configure and deploy the app to a pre-existing Kubernetes cluster

### Built With

* Kustomize

### Prerequisites

    Kubernetes: You need version 1.14 to use Kustomize as a built in tool, otherwhise you will need to download it as a standalone binary
    Kubernetes: You need a running kubernetes cluster with kubectl and/or kustomize binaries installed
    Kubernetes: You need kubectl command configured to talk to your master/api server

### Configuration

#### Background

I am running kubernetes on a rasberry pi cluster which only supports images based on arm architectures, recommend using `cloud`. Also this non-cloud setup does not support cloud load balancers so the type is `NodePort`

1. Choose an environment in overlays `picluster` or `cloud`
2. Configure various yaml files to your liking in overlays/ `deployment_env.yaml,deployment_replicas.yaml,ingress.yaml,kustomization.yaml,namespace.yaml,service.yaml`

### Initialisation

#### Pure kubernetes 

From 1.14 Kustomize is included as part of kubernetes binaries. To run the stack using pure kubernetes run: `kubectl apply -k overlays/picluster` or `kubectl apply -k overlays/cloud`

#### Kustomize standalone

If you have kustomize installed as a standalone binary you need to pipe the output into a kubectl apply command, example: `kustomize build overlays/picluster | kubectl apply -f -` 
    
### Layout

Basic folder layout

    .
    ├── base
    └── overlays
        ├── cloud
        │   └── assets
        ├── custom-metadata-labels
        └── picluster
            └── assets
