package nodepool

import (
	"bytes"
	"fmt"
	"os/exec"
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"os"
	"time"
)

func CreateNodePool(clusterName string, nodePoolName string, rgroupName string, npNodeCount string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Creating nodepool : "+nodePoolName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "aks", "nodepool", "add", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName, "--node-count", npNodeCount)
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

func DeleteNodePool(clusterName string, nodePoolName string, rgroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Deleting nodepool : "+nodePoolName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "aks", "nodepool", "delete", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName)
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
	color.Green("Nodepool Deleted")
}

func ScaleNodePool(clusterName string, nodePoolName string, rgroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Scaling nodepool : "+nodePoolName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "aks", "nodepool", "scale", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName)
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

func UpdateNodePool(clusterName string, nodePoolName string, rgroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Updating nodepool : "+nodePoolName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "aks", "nodepool", "update", "--cluster-name", clusterName, "--name",
		nodePoolName, "--resource-group", rgroupName)
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

func GetNodePool(clusterName string, rgroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Fetching nodepools")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "aks", "nodepool", "list", "--cluster-name", clusterName, "--resource-group", rgroupName)
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
