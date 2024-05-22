/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve system info record(s) from parameter store",
	Long:  `Retrieve system info record(s) from parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&ProfileName, "profile", "p", "", "Profile name used for secrets account")
	getCmd.Flags().StringVarP(&Region, "region", "r", "us-east-1", "Region used for secrets account")
	getCmd.Flags().StringVarP(&Client, "client", "c", "", "Name of client account. Use the client code always!")
	getCmd.Flags().StringVarP(&Environment, "environment", "e", "", "Environment")
	getCmd.Flags().StringVar(&System, "system", "", "System name")
}
