/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ppreeper/odootools/pkg"
	"github.com/spf13/cobra"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "odootools",
	Short: "odootools",
	Long:  `odootools`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Config File
	userConfigDir, err := os.UserConfigDir()
	pkg.CheckErr(err)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", userConfigDir+"/odooquery/config.yml", "odoo connection config file")
}
