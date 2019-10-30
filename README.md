# aksctl
 `aksctl` is an easy to use CLI(Command Line Utility) for creating and managing Kubernetes Cluster on Azure AKS ( Azure Kubernetes Service). It is written in Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

You will need to have Azure credentials configured. What works for AZ CLI should be sufficient. If you dont have Azure CLI installed, you can refer [here](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)


### Installation



## Basic usage
A default cluster can be created by running:
```bash
aksctl create cluster
```
The cluster will be created with default parameters.

Example:

Running the code 
```bash
aksctl create resourcegroup -n myresourcegroup -r eastus
```
creates a resource group with the name `myresourcegroup` and region set to `eastus`.

You can also customize your cluster by using a config file. 
For this, create a `.yaml` file like cluster.yaml
Next, run the command:
```bash
eksctl create cluster -f cluster.yaml
to apply the cluster.yaml file
```
This will create a cluster as described.

Install `aksctl` following the installation instructions.

To learn more abut what aksctl can do check [aksctl.io](https://www.aksctl.com). A good place to start is Getting Started. The full list of features can be found here.


## Deployment



## Built With
* Go