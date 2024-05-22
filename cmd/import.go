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

func MustHave() bool {
	if len(Client) == 0 || len(Environment) == 0 {
		return false
	}
	return true
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import system info records into parameter store",
	Long:  `Import system info records into parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		if !MustHave() {
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
	importCmd.Flags().StringVarP(&ProfileName, "profile", "p", "", "Profile name used for secrets account")
	importCmd.Flags().StringVarP(&Region, "region", "r", "us-east-1", "Region used for secrets account")
	importCmd.Flags().StringVarP(&Client, "client", "c", "", "Name of client account. Use the client code always!")
	importCmd.Flags().StringVarP(&Environment, "environment", "e", "", "Environment")
	importCmd.Flags().StringVar(&System, "system", "", "System name")
}
