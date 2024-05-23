/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/BlueSageSolutions/sysinfo/pkg/sysinfo"
	"github.com/spf13/cobra"
)

func MustHave(strings ...string) bool {
	for _, str := range strings {
		if len(str) == 0 {
			return false
		}
	}
	return true
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import system info records into parameter store",
	Long:  `Import system info records into parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		if !MustHave(Client, Environment) {
			fmt.Println("Missing client or environment")
			os.Exit(1)
		}
		err := sysinfo.MigrateSystemInfo(ProfileName, Region, Client, Environment, System)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVar(&ProfileName, "profile", "", "Profile name used for secrets account")
	importCmd.Flags().StringVar(&Region, "region", "us-east-1", "Region used for secrets account")
	importCmd.Flags().StringVar(&Client, "client", "", "Name of client account. Use the client code always!")
	importCmd.Flags().StringVar(&Environment, "environment", "", "Environment")
	importCmd.Flags().StringVar(&System, "system", "", "System name")
}
