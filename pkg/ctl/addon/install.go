package addon

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/kyokomi/emoji"
)

//InstallAddon will install an addon
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
