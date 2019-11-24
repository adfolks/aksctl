# aksctl

 `aksctl` is an easy to use CLI(Command Line Utility) for creating and managing Kubernetes Cluster on Azure AKS ( Azure Kubernetes Service). It is written in Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

You will need to have Azure credentials configured. What works for AZ CLI should be sufficient. If you dont have Azure CLI installed, you can refer [here](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)


### Installation

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

### Basic usage

#### Cluster

A default cluster can be created by running:
  ```bash
     aksctl create cluster
  ```
  This command will be using values from `default.yaml` situated in root directory,
  edit `default.yaml` with referance of keys from `templet.yaml` for specifying more parameters
  A cluster's mandatrory parameters can be overided using falgs while running command.
  ```bash
     aksctl create cluster --name
                           --rgroupname
                           --rgroupregion
  ```
  Example:

  Running the code
  ```bash
     aksctl create cluster --rgroupname myresourcegroup
  ```
  Flags can be used with shorthand name as well
  ```bash
     aksctl create cluster -r myresourcegroup
  ```
  above code will create a resource group with the name `myresourcegroup`.
  You can also customize your cluster by using a config file.
  For this, create a `.yaml` file like `cluster.yaml` with referance of keys from `templet.yaml`.
  Next, run the command:
  ```bash
     aksctl create cluster -file cluster.yaml
  ```
  to apply the `cluster.yaml` file.
  This will create a cluster as described in the file.

A cluster can be deleted by running:
  ```bash
     aksctl delete cluster
  ```
  A cluster with parameters will be deleted.
  ```bash
      aksctl delete cluster --name --rgroupname
  ```
  Example:

  Running the code
  ```bash
      aksctl delete cluster --rgroupname myresourcegroup
  ```
  above code will delete a resource group with the name `myresourcegroup`.

A cluster can be updated by running:
  ```bash
      aksctl update cluster
  ```
  A cluster with parameters will be updated using flag values.
  ```bash
     aksctl update cluster --name
                           --rgroupname
                           --rgroupregion
  ```
  Example:

  Running the code
  ```bash
     aksctl update cluster --rgroupname myresourcegroup
  ```
  above code will update a resource group with the name `myresourcegroup`.

Get the list of clusters by running:
  ```bash
     aksctl get cluster
  ```
  This gives the list of available clusters.
  ```bash
     aksctl get cluster --rgroupname 
  ```
  Example:

  Running the code
  ```bash
     aksctl get cluster --rgroupname myresourcegroup
  ```

#### Disk

A disk can be created by running:
  ```bash
     aksctl create disk
  ```
  A disk with default parameters will be created.
  ```bash
     aksctl create disk --name 
                        --rgroupname 
                        --rgroupregion
                        --size
  ```  
  Example:

  Running the code
  ```bash
     aksctl create disk --rgroupname myresourceroup
  ```
  above code will create a resource group with the name `myresourcegroup`.

A disk can be deleted by running:
  ```bash
     aksctl delete disk
  ``` 
  A disk with parameters will be deleted.
  ```bash
     aksctl delete disk --name --rgroupname
  ```
  Example:

  Running the code
  ```bash
     aksctl delete disk --rgroupname myresourcegroup
  ```
  above code will delete a resource group with the name `myresourcegroup`.

A disk can be updated by running:
  ```bash
     aksctl update disk
  ```
  A disk with parameters will be updated using flag values.      
  ```bash
    aksctl update disk --name
                       --rgroupname
                       --size                            
  ```
  Example:

  Running the code
  ```bash
     aksctl update disk --rgoupname myresourcegroup        
  ```
  above code will update a resource group with the name `myresourcegroup`.

Get the list of disks by running:
  ```bash
      aksctl get disk
  ```
  This gives the list of available disks.
  ```bash
     aksctl get disk --rgroupname
  ```
  Example:

  Running the code
  ```bash
     aksctl get disk --rgroupname myresourcegroup
  ```

#### Nodepool

A nodepool can be created by running:
  ```bash
     aksctl create nodepool
  ```
  A nodepool with default parameters will be created.
  ```bash
     aksctl create nodepool --nodepoolname
                            --rgroupname
                            --clustername
  ```
  Example:

  Running the code            
  ```bash
     aksctl create nodepool --nodepoolname mynode
  ```
  above code will create a nodepool with the name `mynode`.

A nodepool can be deleted by running:
   ```bash
      aksctl delete nodepool
   ```
   A nodepool with parameters will be deleted.
   ```bash
      aksctl delete nodepool --nodepoolname 
                             --rgroupname
                             --clustername
   ```
   Example:

   Running the code
   ```bash
      aksctl delete nodepool --nodepoolname mynode
   ```
   above code will delete a nodepool with the name `mynode`.

A nodepool can be updated by running:
   ```bash
      aksctl update nodepool
   ```
   A nodepool with parameters will be updated using flag values.
  
   ```bash
      aksctl update nodepool --nodepoolname
                             --rgroupname
                             --clustername
   ```
   Example:

   Running the code
   ```bash
      aksctl update nodepool --nodepoolname mynode
   ```
   above code will update a nodepool with the name `mynode`.

A nodepool can be scaled by running:
   ```bash
      aksctl scale nodepool
   ```
   Using scale, we can scale nodepool in a kubernetes cluster.
   ```bash
      aksctl scale nodepool --nodepoolname
                            --rgroupname
                            --clustername
   ```
   Example:

   Running the code
   ```bash
      aksctl scale nodepool --nodepoolname mynode
   ```

Get the list of nodepools by running:
   ```bash
      aksctl get nodepool
   ```
   This gives the list of available nodepools.
   ```bash
      aksctl get nodepool --rgroupname --clustername
   ```
   Example:

   Running the code
   ```bash
      aksctl get nodepool --rgroupname myresourcegroup
   ```

#### ResorceGroup

A resource group can be created by running:
   ```bash
      aksctl create resource group
   ```
   A resource group with default parameters will be created.
   ```bash
      aksctl create resourcegroup --rgroupname --rgroupregion
   ```  
   Example:

   Running the code
   ```bash
      aksctl create resourcegroup --rgroupname myresourcegroup
   ```
   above code will create a resource group with the name `myresourcegroup`.

A resource group can be deleted by running:
  ```bash
      aksctl delete resource group
  ``` 
  A resource group with parameters will be deleted.
  ```bash
     aksctl delete resourcegroup --rgroupname
  ```
  Example:

  Running the code
  ```bash
     aksctl delete resourcegroup --rgroupname myresourcegroup
  ```
  above code will delete a resource group with the name `myresourcegroup`.
  
A resource group can be updated by running:
  ```bash
     aksctl update resourcegroup
  ```
  A resourcegroup with parameters will be updated using flag values.      
  ```bash
     aksctl update resourcegroup --rgroupname                         
  ```
  Example:

  Running the code
   ```bash
      aksctl update resourcegroup --rgroupname myresourcegroup        
   ```
   above code will update a resourcegroup with the name `myresourcegroup`.

Get the list of resource groups by running:
   ```bash
      aksctl get resource group
   ```
   This gives the list of available resource groups.

#### Vnet

A vnet can be created by running:
  ```bash
     aksctl create vnet
  ```
  A vnet with default parameters will be created.
  ```bash
     aksctl create vnet --vnetname --rgroupname
  ```
  Example:

  Running the code            
  ```bash
     aksctl create vnet --rgroupname myresourcegroup
  ```
  above code will create a resourcegroup with the name `myresourcegroup`.

A vnet can be deleted by running:
   ```bash
      aksctl delete vnet
   ```
   A vnet with parameters will be deleted.
   ```bash
      aksctl delete vnet --vnetname --rgroupname
   ```
   Example:

   Running the code
   ```bash
      aksctl delete vnet --rgroupname myresourcegroup
   ```
   above code will delete a resourcegroup with the name `myresourcegroup`.
 
`aksctl` can be installed by following the above instructions.

Check [aksctl.io](https://www.aksctl.com) to learn more abut what aksctl can do and its features.

## Built With
* Go