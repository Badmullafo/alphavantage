# alphavantage

<!-- ABOUT THE PROJECT -->
## About The Project

This project provides resources and instructions to build a simple python based web scraping app. You can deploy it to a Kubernetes cluster using Kustomize

More detailed README.md files and instructions can be found within subfolders

### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* Docker
* Kustomize

### Prerequisites

    Kubernetes: you need version 1.14 to use Kustomize as a built in tool, otherwhise you will need to download it as a standalone binary
    Git: You need git installed to clone this repo
    Docker: You need docker installed if you wish to build images from dockerfile
    
### Layout

Basic folder layout

    .
    ├── kustomize
    │   ├── base
    │   └── overlays
    │       ├── custom-metadata-labels
    │       ├── dev
    │       └── prod
    │           └── assets
    └── python_web
