<!-- ABOUT THE PROJECT -->
## About

This directory contains instructions and some artifacts to build the docker image to run the app

### Built With

* Docker

### Prerequisites

    Docker: You need docker installed if you wish to build images from dockerfile

### Installation - Build with docker (OPTIONAL)

If you want to build your own docker image `docker build -t yourepo/alphavantage:1.0 .`

**Warning, where you build this image affects which architecture it is based on**

You can push the image to a private registry if you want, just make sure you tag it `-t` correctly so you can push it to private image registry

### Get from dockerhub  (OPTIONAL)

    docker pull badmullafo/alphavantage:1.0

Alternatively, if you do not want/need to build the image you can get it from dockerhub. You don't need to worry about this, Kustomize will handle this for you (See Kustomize section)

### Limitations

The web content is static not dynamic - it will not get updated
