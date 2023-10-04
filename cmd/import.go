/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import data into an Odoo database",
	Long:  `import data into an Odoo database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVar(&sHost.Hostname, "hostname", "", "odoo hostname")
	importCmd.Flags().StringVar(&sHost.Database, "database", "", "odoo database")
	importCmd.Flags().StringVar(&sHost.Username, "username", "", "odoo username")
	importCmd.Flags().StringVar(&sHost.Password, "password", "", "odoo password")
	importCmd.Flags().StringVar(&sHost.Schema, "schema", "http", "odoo url schema [http|https]")
	importCmd.Flags().IntVar(&sHost.Port, "port", 8069, "odoo port")

	importCmd.Flags().StringVarP(&q.Model, "model", "m", "", "model")
	importCmd.Flags().StringVar(&q.Filter, "filter", "", "filter")
	importCmd.Flags().IntVar(&q.Offset, "offset", 0, "offset")
	importCmd.Flags().IntVar(&q.Limit, "limit", 0, "limit")
	importCmd.Flags().StringVar(&q.Fields, "fields", "", "fields")
	importCmd.Flags().BoolVar(&q.Count, "count", false, "count")
	importCmd.MarkFlagRequired("model")
}
