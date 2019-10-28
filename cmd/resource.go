package cmd

import (
	"fmt"
	"github.com/Roshni1313/aksctl/coreaksctl"
	"github.com/spf13/cobra"
)

var rgroupName, rgroupRegion string

var resourcegroupCmd = &cobra.Command{
	Use:   "resourcegroup",
	Short: "Create an AKS resource group",
	Long: `Create an AKS resource group, it would use a Random Name for resource group.
	If you need to specify name or other resources use resourcegroup.yaml file for more custom configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		rgroupName, _ = cmd.Flags().GetString("name")
		rgroupRegion, _ = cmd.Flags().GetString("region")
		coreaksctl.CreateResourceGroup(rgroupName, rgroupRegion)
		fmt.Println("Resource group created")
	},
}

func init() {
	createCmd.AddCommand(resourcegroupCmd)
	resourcegroupCmd.PersistentFlags().StringP("name", "n", "temp", "resource group name")
	resourcegroupCmd.PersistentFlags().StringP("region", "r", "westus", "resource group location")
}
