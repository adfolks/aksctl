package addon

import (
	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"os"
	"time"
	"github.com/kyokomi/emoji"
)

func InstallAddon() {
	a := wow.New(os.Stdout, spin.Get(spin.Dots), "Installing your addon")
	a.Start()
	time.Sleep(2 * time.Second)
	a.Text("This would take a few minutes...").Spinner(spin.Get(spin.Dots))
	//Install Addon code here
	
	a.PersistWith(spin.Spinner{}, "....")
	color.Green("Addon installed")
    emoji.Println(":beer: Cheers!!!")
}
