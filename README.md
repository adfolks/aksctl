# aksctl

 `aksctl` is an easy to use CLI(Command Line Utility) for creating and managing Kubernetes Cluster on Azure AKS ( Azure Kubernetes Service). It is written in Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

You will need to have Azure credentials configured. What works for AZ CLI should be sufficient. If you dont have Azure CLI installed, you can refer [here](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)


### Installation

#### Install from source
To install aksctl, clone the repository using:
```bash
git clone https://github.com/adfolks/aksctl
```
Then, initiate module with:
```bash
go mod init github.com/adfolks/aksctl
```
This will create a module config file `go.mod`.
Finally, use
```bash
go build
```
to fetch the latest dependencies.

## Basic usage
A default cluster can be created by running:
```bash
aksctl create cluster
```
A cluster with default parameters will be created.

Example:

Running the code 
```bash
aksctl create resourcegroup --name myresourcegroup --region eastus
```
creates a resource group with the name `myresourcegroup` and region set to `eastus`.

You can also customize your cluster by using a config file. 
For this, create a `.yaml` file like `cluster.yaml`.

Next, run the command:
```bash
aksctl create cluster -f cluster.yaml
```
to apply the `cluster.yaml` file.
This will create a cluster as described in the file.

`aksctl` can be installed by following the above instructions.

Check [aksctl.io](https://www.aksctl.io) to learn more abut what aksctl can do and its features.

## Built With
* Go