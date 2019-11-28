package disk

import (
	"bytes"
	"fmt"
	"os/exec"

	"os"
	"time"

	"github.com/adfolks/aksctl/pkg/ctl/utils"
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/kyokomi/emoji"
)

//CreateDisk will create a disk
func CreateDisk(diskName string, diskResourcegroup string, diskLocation string, diskSize string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Creating disk : "+diskName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Create AKS Diska.PersistWith(spin.Spinner{}, "....")
	cmd := exec.Command("az", "disk", "create", "-g", diskResourcegroup, "-n",
		diskName, "-l", diskLocation, "-z", diskSize)
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
	color.Green("Disk Created")
	emoji.Println(":beer: Cheers!!!")
}

//DeleteDisk will delete a disk
func DeleteDisk(diskName string, diskResourceGroup string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Deleting disk : "+diskName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Delete AKS Disk
	cmd := exec.Command("az", "disk", "delete", "--name", diskName,
		"--resource-group", diskResourceGroup, "--yes")
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
	color.Green("Disk Deleted")
}

//UpdateDisk will modify the changes
func UpdateDisk(diskName string, diskResourceGroup string, diskSize string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Updating disk : "+diskName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Update AKS Disk
	cmd := exec.Command("az", "disk", "update", "--name", diskName,
		"--resource-group", diskResourceGroup, "--size-gb", diskSize)
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
	color.Green("Disk Updated")
}

//GetDisk will list the available disks
func GetDisk(diskResourceGroup string, diskFilter string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Fetching disks")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//List AKS Disk
	cmd := exec.Command("az", "disk", "list",
		"--resource-group", diskResourceGroup)
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
	if diskFilter == "all" {
		fmt.Println("Result: " + out.String())
	} else {
		fmt.Println("Result:", utils.FilterStringMap(out.String(), diskFilter))
	}
}
