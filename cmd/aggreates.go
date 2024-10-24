package cmd

import (
	"github.com/AlissonBarbosa/alfred/controllers"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(aggregateCmd)
	aggregateCmd.Flags().StringP("os-cloud", "", "default", "Specify the cloud from clouds.yaml")
}

var aggregateCmd = &cobra.Command{
	Use:   "aggregate list",
	Short: "List aggregates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := args[0]
		if status != "list" {
			log.Fatalf("Invalid usage. Use list.")
		}
		cloud, err := cmd.Flags().GetString("os-cloud")
		if err != nil {
			log.Fatalf("Could not get cloud: %v", err)
		}
		//controllers.ListAggregates(cloud)
		headers, columns, err := controllers.ListAggregates(cloud)
		if err != nil {
			log.Fatalf("Error getting list aggregates: %v", err)
		}
		controllers.RenderTable(headers, columns)
	},
}
