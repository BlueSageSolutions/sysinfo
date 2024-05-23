/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	secrets "github.com/BlueSageSolutions/secrets/cmd"
	"github.com/BlueSageSolutions/sysinfo/pkg/sysinfo"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve system info record(s) from parameter store",
	Long:  `Retrieve system info record(s) from parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		systemInfo, err := secrets.GetSecret("", ProfileName, Region, "secrets", Client, Environment, sysinfo.Sluggify(System), "system-info", "", true, true)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			os.Exit(1)
		}
		fmt.Println(systemInfo)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVar(&ProfileName, "profile", "", "Profile name used for secrets account")
	getCmd.Flags().StringVar(&Region, "region", "us-east-1", "Region used for secrets account")
	getCmd.Flags().StringVar(&Client, "client", "", "Name of client account. Use the client code always!")
	getCmd.Flags().StringVar(&Environment, "environment", "", "Environment")
	getCmd.Flags().StringVar(&System, "system", "", "System name")
}
