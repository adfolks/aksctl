package addon

import (
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"os"
	"time"
	"bytes"
	"os/exec"
	"github.com/kyokomi/emoji"
	"fmt"
	"strings"
)

func InstallAddon(chartName string, repoName string) {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Installing your addon")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Install Addon code here
	cmd := exec.Command("helm", "install", chartName, repoName)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	
	a.PersistWith(spin.Spinner{}, "....")
	color.Green("Addon installed")
    emoji.Println(":beer: Cheers!!!")
}

func CheckHelm()bool{
	fmt.Println("Checking if latest version of helm is present")
	cmd := exec.Command("helm", "version")
	versioncheck := false
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Install helm to use charts")
	}else{
	versioncheck = strings.Contains(out.String(), "Version:\"v3.")
	}
	return versioncheck
}
