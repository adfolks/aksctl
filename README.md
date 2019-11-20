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
1.A default cluster can be created by running:
  ```bash
     aksctl create cluster
  ```
   A cluster with default parameters will be created.
  ```bash
     aksctl create --name
                   --rgroupname
                   --rgroupregion
  ```
  Example:
  Running the code            
  ```bash
     aksctl create --rgroupname myresourcegroup 
  ```
   creates a resource group with the name `myresourcegroup`.

   You can also customize your cluster by using a config file. 
   For this, create a `.yaml` file like `cluster.yaml`.

   Next, run the command:
  ```bash
     aksctl create cluster -f cluster.yaml
   ```
   to apply the `cluster.yaml` file.
   This will create a cluster as described in the file.
  
2.A cluster can be deleted by running:
   ```bash
      aksctl delete cluster
   ```
   A cluster with parameters will be deleted.
   ```bash
      aksctl delete --name --rgroupname
   ```
  Example:
  Running the code
   ```bash
      aksctl delete --name myclustername 
   ```
   deletes a cluster with the name `myclustername`.
  
3.A cluster can be updated by running:
   ```bash
      aksctl update cluster
   ```
  A cluster with parameters will be updated using flag values.
  
  ```bash
     aksctl update --name
                   --rgroupname
                   --rgroupregion
  ```
  Example:
  Running the code
  ```bash
     aksctl update --name myclustername 
  ```
  updates a cluster with the name `myclustername`.

4.Get the list of clusters by running:
  ```bash
     aksctl get cluster
  ```
  This gives the list of available clusters.
  ```bash
     aksctl get --rgroupname 
  ```
  Example:
  Running the code
  ```bash
     aksctl get --rgroupname myresourcegroup
  ```

5.A disk can be created by running:
  ```bash
     aksctl create disk
  ```
  A disk with default parameters will be created.
  ```bash
     aksctl create --name 
                   --rgroupname 
                   --rgroupregion
                   --size
  ```  
  Example:
  Running the code
  ```bash
     aksctl create --name mydisk 
  ```
  creates a disk with the name `mydisk`.

6.A disk can be deleted by running:

  ```bash
     aksctl delete disk
  ``` 
  A disk with parameters will be deleted.
  ```bash
     aksctl delete --name --rgroupname
  ```
  Example:
  Running the code
  ```bash
     aksctl delete --name mydisk 
  ```
  deletes a disk with the name `mydisk`.
  
7.A disk can be updated by running:
  ```bash
     aksctl update disk
  ```
  A disk with parameters will be updated using flag values.      
  ```bash
    aksctl update --name
                  --rgroupname
                  --size                            
  ```
  Example:
  Running the code
   ```bash
     aksctl update --name mydisk 
             
   ```
   updates a disk with the name `mydisk`.

8.Get the list of disks by running:
   ```bash
      aksctl get disk
   ```
   This gives the list of available disks.
   ```bash
      aksctl get --rgroupname
   ```
   Example:
   Running the code
   ```bash
     aksctl get --rgroupname myresourcegroup
   ```

9.A nodepool can be created by running:
  ```bash
     aksctl create nodepool
  ```
   A nodepool with default parameters will be created.
  ```bash
     aksctl create --nodepoolname
                   --rgroupname
                   --clustername
  ```
  Example:
  Running the code            
  ```bash
     aksctl create --nodepoolname mynode
  ```
   creates a nodepool with the name `mynode`.

10.A nodepool can be deleted by running:
   ```bash
      aksctl delete nodepool
   ```
   A nodepool with parameters will be deleted.
   ```bash
      aksctl delete --nodepoolname 
                    --rgroupname
                    --clustername
   ```
  Example:
  Running the code
   ```bash
      aksctl delete --nodepoolname mynode
   ```
   deletes a nodepool with the name `mynode`.

11.A nodepool can be updated by running:
   ```bash
      aksctl update nodepool
   ```
  A nodepool with parameters will be updated using flag values.
  
  ```bash
     aksctl update --nodepoolname
                   --rgroupname
                   --clustername
  ```
  Example:
  Running the code
  ```bash
     aksctl update --nodepoolname mynode
  ```
  updates a nodepool with the name `mynode`.

12.A nodepool can be scaled by running:
   ```bash
      aksctl scale nodepool
   ```
   Using scale, we can scale nodepool in a kubernetes cluster.
   ```bash
      aksctl scale --nodepoolname
                   --rgroupname
                   --clustername
   ```
   Example:
   Running the code
   ```bash
      aksctl scale --nodepoolname mynode
   ```

13.Get the list of nodepools by running:
  ```bash
     aksctl get nodepool
  ```
  This gives the list of available nodepools.
  ```bash
     aksctl get --rgroupname --clustername
  ```
  Example:
  Running the code
  ```bash
     aksctl get --rgroupname myresourcegroup
  ```

14.A resource group can be created by running:
  ```bash
     aksctl create resource group
  ```
  A resource group with default parameters will be created.
  ```bash
     aksctl create --rgroupname --rgroupregion
  ```  
  Example:
  Running the code
  ```bash
     aksctl create --rgroupname myresourcegroup
  ```
  creates a resource group with the name `myresourcegroup`.

15.A resource group can be deleted by running:

  ```bash
     aksctl delete resource group
  ``` 
  A resource group with parameters will be deleted.
  ```bash
     aksctl delete --rgroupname
  ```
  Example:
  Running the code
  ```bash
     aksctl delete --rgroupname myresourcegroup
  ```
  deletes a resource group with the name `myresourcegroup`.
  
16.A resource group can be updated by running:
  ```bash
     aksctl update resourcegroup
  ```
  A resourcegroup with parameters will be updated using flag values.      
  ```bash
    aksctl update --rgroupname                         
  ```
  Example:
  Running the code
   ```bash
     aksctl update --rgroupname myresourcegroup
             
   ```
   updates a resourcegroup with the name `myresourcegroup`.

17.Get the list of resource groups by running:
   ```bash
      aksctl get resource group
   ```
   This gives the list of available resource groups.
   
`aksctl` can be installed by following the above instructions.

Check [aksctl.io](https://www.aksctl.com) to learn more abut what aksctl can do and its features.


## Built With
* Go