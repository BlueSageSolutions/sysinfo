/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export system info records from parameter store",
	Long:  `Export system info records from parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&ProfileName, "profile", "p", "", "Profile name used for secrets account")
	exportCmd.Flags().StringVarP(&Region, "region", "r", "us-east-1", "Region used for secrets account")
	exportCmd.Flags().StringVarP(&Client, "client", "c", "", "Name of client account. Use the client code always!")
	exportCmd.Flags().StringVarP(&Environment, "environment", "e", "", "Environment")
	exportCmd.Flags().StringVar(&System, "system", "", "System name")

}
