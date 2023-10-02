/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ppreeper/odootools/pkg"

	"github.com/spf13/cobra"
)

var (
	cServer pkg.Host
	q       pkg.QueryDef
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
		if server.Hostname == "" && cServer.Hostname == "" {
			fmt.Println("no host specified")
			return
		}
		if cServer.Hostname != "" {
			server.Hostname = cServer.Hostname
		}

		// 	Database
		if server.Database == "" && cServer.Database == "" {
			fmt.Println("no database specified")
			return
		}
		if cServer.Database != "" {
			server.Database = cServer.Database
		}

		// 	Username
		if cServer.Username != "" {
			server.Username = cServer.Username
		}
		// 	Password
		if cServer.Password != "" {
			server.Password = cServer.Password
		}
		// 	Protocol
		if server.Protocol == "" && cServer.Protocol != "" {
			server.Protocol = cServer.Protocol
		}
		// 	Schema
		if server.Schema == "" && cServer.Schema != "" {
			server.Schema = cServer.Schema
		}
		// 	Port
		if server.Port == 0 {
			server.Port = cServer.Port
		}
		// 	Jobcount
		if server.Jobcount == 0 {
			server.Jobcount = cServer.Jobcount
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

	queryCmd.Flags().StringVar(&cServer.Hostname, "hostname", "", "odoo hostname")
	queryCmd.Flags().StringVar(&cServer.Database, "database", "", "odoo database")
	queryCmd.Flags().StringVar(&cServer.Username, "username", "", "odoo username")
	queryCmd.Flags().StringVar(&cServer.Password, "password", "", "odoo password")
	queryCmd.Flags().StringVar(&cServer.Schema, "schema", "http", "odoo url schema [http|https]")
	queryCmd.Flags().IntVar(&cServer.Port, "port", 8069, "odoo port")

	queryCmd.Flags().StringVarP(&q.Model, "model", "m", "", "model")
	queryCmd.Flags().StringVar(&q.Filter, "filter", "", "filter")
	queryCmd.Flags().IntVar(&q.Offset, "offset", 0, "offset")
	queryCmd.Flags().IntVar(&q.Limit, "limit", 0, "limit")
	queryCmd.Flags().StringVar(&q.Fields, "fields", "", "fields")
	queryCmd.Flags().BoolVar(&q.Count, "count", false, "count")
	queryCmd.MarkFlagRequired("model")
}
