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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a system info record in parameter store",
	Long:  `Create a system info record in parameter store`,
	Run: func(cmd *cobra.Command, args []string) {
		if !MustHave(Client, Environment, System) {
			fmt.Println("Check usage")
			os.Exit(1)
		}
		value := sysinfo.Denormalize(System, URL, Username, Password, Notes, SystemProperties, ConfigData, ValidationClass, AwsSystemType, Enabled)

		err := sysinfo.SetSystemInfo(ProfileName, Region, sysinfo.Sluggify(System), Client, Environment, value)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVar(&ProfileName, "profile", "", "Profile name used for secrets account")
	setCmd.Flags().StringVar(&Region, "region", "us-east-1", "Region used for secrets account")
	setCmd.Flags().StringVar(&Client, "client", "", "Name of client account. Use the client code always!")
	setCmd.Flags().StringVar(&Environment, "environment", "", "Environment")
	setCmd.Flags().StringVar(&System, "system", "", "System name")
	setCmd.Flags().StringVar(&URL, "url", "", "URL")
	setCmd.Flags().StringVar(&Username, "username", "", "Username")
	setCmd.Flags().StringVar(&Password, "password", "", "Password")
	setCmd.Flags().StringVar(&Notes, "notes", "", "Notes")
	setCmd.Flags().StringVar(&SystemProperties, "system-properties", "", "System Properties")
	setCmd.Flags().StringVar(&ConfigData, "config-data", "", "Config Data")
	setCmd.Flags().StringVar(&ValidationClass, "validation-class", "", "Validation Class")
	setCmd.Flags().StringVar(&AwsSystemType, "aws-system-type", "", "Aws System Type")
	setCmd.Flags().BoolVar(&Enabled, "enabled", false, "Enabled")

}
