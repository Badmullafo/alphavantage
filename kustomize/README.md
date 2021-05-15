# kustomize

<!-- ABOUT THE PROJECT -->
## About The Project

This readme provides instructions on how to configure and deploy the app to a pre-existing Kubernetes cluster. The instructions use Kustomise to template manifests

### Built With

* Kustomize

### Prerequisites

    Kubernetes: You need version 1.14 to use Kustomize as a built in tool, otherwhise you will need to download it as a standalone binary
    Kubernetes: You need a running kubernetes cluster with kubectl and/or kustomize binaries installed
    Kubernetes: You need kubectl command configured to talk to your master/api server

### Configuration

#### Background

I am running kubernetes on a rasberry pi cluster which only supports images based on arm architectures, recommend using `cloud`. Also this non-cloud setup does not support cloud load balancers so the type is `NodePort`. I tried implementing a kind of virtual load balancer using [metallb](https://metallb.universe.tf/installation/) there are certain limitations to the [layer2](https://metallb.universe.tf/concepts/layer2/#limitations) setup, I found it to be unrealiable

1. Choose an environment in overlays  `cloud` or `picluster` - **recommend cloud as picluster is arm based and won't work on non arm**
2. Configure various yaml files to your liking in `overlays/` `deployment_env.yaml,deployment_replicas.yaml,ingress.yaml,kustomization.yaml,namespace.yaml,service.yaml`

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
        └── picluster

### Testing

To test it has been deployed, go to the web page, use the `/stock` appended URL or whatever you specifieed in your `ingress.yaml`

    http://myapp.com/stock

That should display the result, if you have more than on replica, if you refresh the web page repeatedly you should get a different pod serve you the page each time, for example the host will change (see top line)

    Request: /stock from host alphavantage-5664d6c6bf-4xgd7

    Getting the last 10 days worth of results for IBM, the list is [141.3, 144.22, 146.17, 145.46, 148.42, 145.22, 145.75, 144.75] average is 145.16.

### Limitations

* Would recommend not upping the replica count to more than 4 as the API key seems to have some kind of limitation on how often it can be used in a short space of time
* Only works as nodeport on picluster
