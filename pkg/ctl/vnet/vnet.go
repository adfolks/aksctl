package vnet

import (
	"bytes"
	"fmt"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/kyokomi/emoji"
	"os"
	"os/exec"
	"time"
)

//CreatVNet will create vnet
func CreateVNet(vNetName string, rgroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Creating virtual network : "+vNetName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "network", "vnet", "create", "--name",
		vNetName, "--resource-group", rgroupName)
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
	emoji.Println(":beer: Cheers!!!")
}

//DeleteVNet will delete vnet
func DeleteVNet(vNetName string, rgroupName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Deleting virtual network : "+vNetName)
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	cmd := exec.Command("az", "network", "vnet", "delete", "--name",
		vNetName, "--resource-group", rgroupName)
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
