package resourcegroup

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"os"
	"time"
	"os/exec"
)

func CreateResourceGroup(rgroupName string, rgroupRegion string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Creating resource group : "+rgroupName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "group", "create", "-l", rgroupRegion, "-n", rgroupName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	a.PersistWith(spin.Spinner{}, "....")
	color.Green("Resource group Created")

}

func CheckResourceGroup(rgroupName string) bool {

	cmd := exec.Command("az", "group", "exists", "-n", rgroupName)
	var check bool
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	if out.String() == "true\n" {
		check = true
		fmt.Println("Using existing resource group")
	} else {
		check = false
		color.Red("The resource group "+ rgroupName +" does not exist.")
	}
	return check
}

func DeleteResourceGroup(rgroupName string) {

	//Delete AKS ResourceGroup
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Deleting resource group : "+rgroupName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "group", "delete", "--name", rgroupName, "--yes")
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
	color.Green("Resource group Deleted")
}

func UpdateResourceGroup(rgroupName string) {

	//Update AKS ResourceGroup
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Updating resource group : "+rgroupName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))

	cmd := exec.Command("az", "group", "update", "--name", rgroupName)
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
	color.Green("Resource Group Updated")
}

func GetResourceGroup() {

	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Fetching resource groups")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//List AKS ResourceGroup
	cmd := exec.Command("az", "group", "list")
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
