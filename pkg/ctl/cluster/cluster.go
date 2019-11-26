package cluster

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/kyokomi/emoji"
	"os"
	"os/exec"
	"time"
)

func CreateCluster(clusterName string, resourceGroupName string, extraflags []string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Starting to set up your k8s Cluster")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Create AKS Cluster
	var args = []string{"aks", "create", "--name", clusterName, "--resource-group", resourceGroupName}
	//handle the SSH keys generation in case of basic usage or the key value is not present on the config file
	fmt.Println(extraflags)
	_, found := Find(extraflags, "ssh-key-value")
	if !found {
		extraflags = append(extraflags, "--generate-ssh-keys")
	}
	args = append(args, extraflags...)
	cmd := exec.Command("az", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed :" + fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	a.PersistWith(spin.Spinner{}, "....")
	color.Green("Cluster Created")
	GetClusterCredentials(clusterName, resourceGroupName)
	color.Green("Updated Kubeconfig")
	emoji.Println(":beer: Cheers!!!")
}

func GetClusterCredentials(clusterName string, resourceGroupName string) {
	b := wow.New(os.Stdout, spin.Get(spin.Dots), "Fetching credentials")
	b.Start()
	time.Sleep(2 * time.Second)
	b.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Get AKS Cluster credentials
	var args = []string{"aks", "get-credentials", "--overwrite-existing", "--name", clusterName, "--resource-group", resourceGroupName}
	cmd := exec.Command("az", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed :" + fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	b.PersistWith(spin.Spinner{}, "....")
	color.Cyan(out.String())
}

func DeleteCluster(clusterName string, resourceGroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Deleting your k8s Cluster "+clusterName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Delete AKS Cluster
	cmd := exec.Command("az", "aks", "delete", "--name", clusterName,
		"--resource-group", resourceGroupName, "--yes")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	a.PersistWith(spin.Spinner{}, "....")
	color.Green("Cluster Deleted")
}

func UpdateCluster(clusterName string, resourceGroupName string) {

	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Updating your k8s Cluster")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "update", "--name", clusterName,
		"--resource-group", resourceGroupName, "--enable-cluster-autoscaler")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	a.PersistWith(spin.Spinner{}, "....")
	color.Green("Cluster Updated")
}

func GetCluster(resourceGroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Collecting your k8s Cluster informations")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Create AKS Cluster
	cmd := exec.Command("az", "aks", "list",
		"--resource-group", resourceGroupName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	a.PersistWith(spin.Spinner{}, "....")
	fmt.Println("Result: " + out.String())
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
