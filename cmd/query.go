/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ppreeper/odootools/pkg"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:        "query",
	Short:      "query an Odoo database",
	Long:       `query an Odoo database`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"system"},
	Run: func(cmd *cobra.Command, args []string) {
		system := args[0]
		// get config file
		HostMap := pkg.GetConf(configFile)

		// Server connection profile
		server := HostMap[system]
		// fmt.Println("server", server)
		if server.Hostname == "" && sHost.Hostname == "" {
			fmt.Println("no host specified")
			return
		}
		if sHost.Hostname != "" {
			server.Hostname = sHost.Hostname
		}

		// 	Database
		if server.Database == "" && sHost.Database == "" {
			fmt.Println("no database specified")
			return
		}
		if sHost.Database != "" {
			server.Database = sHost.Database
		}

		// 	Username
		if sHost.Username != "" {
			server.Username = sHost.Username
		}
		// 	Password
		if sHost.Password != "" {
			server.Password = sHost.Password
		}
		// 	Protocol
		if server.Protocol == "" && sHost.Protocol != "" {
			server.Protocol = sHost.Protocol
		}
		// 	Schema
		if server.Schema == "" && sHost.Schema != "" {
			server.Schema = sHost.Schema
		}
		// 	Port
		if sHost.Port != 0 {
			server.Port = sHost.Port
		}
		// 	Jobcount
		if sHost.Jobcount != 0 {
			server.Jobcount = sHost.Jobcount
		}

		if q.Model == "" {
			fmt.Println("no model specified")
			return
		}
		oc, err := pkg.OdooConnect(server)
		pkg.FatalErr(err)

		app := pkg.Application{
			OC: oc,
		}
		app.GetRecords(&q)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	queryCmd.Flags().StringVar(&sHost.Hostname, "hostname", "", "odoo hostname")
	queryCmd.Flags().StringVar(&sHost.Database, "database", "", "odoo database")
	queryCmd.Flags().StringVar(&sHost.Username, "username", "", "odoo username")
	queryCmd.Flags().StringVar(&sHost.Password, "password", "", "odoo password")
	queryCmd.Flags().StringVar(&sHost.Schema, "schema", "", "odoo url schema [http|https]")
	queryCmd.Flags().IntVar(&sHost.Port, "port", 0, "odoo port")

	queryCmd.Flags().StringVarP(&q.Model, "model", "m", "", "model")
	queryCmd.Flags().StringVar(&q.Filter, "filter", "", "filter")
	queryCmd.Flags().IntVar(&q.Offset, "offset", 0, "offset")
	queryCmd.Flags().IntVar(&q.Limit, "limit", 0, "limit")
	queryCmd.Flags().StringVar(&q.Fields, "fields", "", "fields")
	queryCmd.Flags().BoolVar(&q.Count, "count", false, "count")
	queryCmd.MarkFlagRequired("model")
}
