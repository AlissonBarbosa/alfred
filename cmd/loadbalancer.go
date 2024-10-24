package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(loadbalancerCmd)
	loadbalancerCmd.Flags().StringP("os-cloud", "", "default", "Specify the cloud from clouds.yaml")
}

var loadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer [ERROR|PENDING_UPDATE]",
	Short: "List loadbalancers by status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := args[0]
		if status != "ERROR" && status != "PENDING_UPDATE" {
			log.Fatalf("Invalid status. Use ERROR or PENDING_UPDATE.")
		}
		cloud, err := cmd.Flags().GetString("os-cloud")
		if err != nil {
			log.Fatalf("Could not get cloud: %v", err)
		}
		fmt.Printf(status, cloud)
	},
}
