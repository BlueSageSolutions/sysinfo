/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a system info record in parameter store",
	Long:  `Create a system info record in parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&ProfileName, "profile", "p", "", "Profile name used for secrets account")
	setCmd.Flags().StringVarP(&Region, "region", "r", "us-east-1", "Region used for secrets account")
	setCmd.Flags().StringVarP(&Client, "client", "c", "", "Name of client account. Use the client code always!")
	setCmd.Flags().StringVarP(&Environment, "environment", "e", "", "Environment")
	setCmd.Flags().StringVar(&System, "system", "", "System name")
	setCmd.Flags().StringVar(&URL, "url", "", "URL")
	setCmd.Flags().StringVar(&UserName, "username", "", "Username")
	setCmd.Flags().StringVar(&Password, "password", "", "Password")
	setCmd.Flags().StringVar(&Notes, "notes", "", "Notes")
	setCmd.Flags().StringVar(&SystemProperties, "system-properties", "", "System Properties")
	setCmd.Flags().StringVar(&ConfigData, "config-data", "", "Config Data")
	setCmd.Flags().StringVar(&ValidationClass, "validation-class", "", "Validation Class")
	setCmd.Flags().StringVar(&AwsSystemType, "aws-system-type", "", "Aws System Type")
	setCmd.Flags().BoolVar(&Enabled, "enabled", false, "Enabled")

}
