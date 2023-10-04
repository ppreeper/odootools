/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/ppreeper/odootools/pkg"
	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:        "copy",
	Short:      "copy data between databases",
	Long:       `copy data between databases`,
	Args:       cobra.MinimumNArgs(2),
	ArgAliases: []string{"source", "dest"},
	Run: func(cmd *cobra.Command, args []string) {
		sSystem := args[0]
		dSystem := args[1]
		if sSystem == dSystem {
			fmt.Println("cannot copy to itself")
			return
		}
		// get config file
		HostMap := pkg.GetConf(configFile)

		// Server connection profile
		sServer := HostMap[sSystem]
		dServer := HostMap[dSystem]
		fmt.Println("sServer", sServer)
		fmt.Println("dServer", dServer)

		if sServer.Hostname == "" && sHost.Hostname == "" {
			fmt.Println("no source host specified")
			return
		}
		if sHost.Hostname != "" {
			sServer.Hostname = sHost.Hostname
		}
		//
		if dServer.Hostname == "" && dHost.Hostname == "" {
			fmt.Println("no destination host specified")
			return
		}
		if dHost.Hostname != "" {
			dServer.Hostname = dHost.Hostname
		}

		// 	Database
		if sServer.Database == "" && sHost.Database == "" {
			fmt.Println("no source database specified")
			return
		}
		if sHost.Database != "" {
			sServer.Database = sHost.Database
		}
		//
		if dServer.Database == "" && dHost.Database == "" {
			fmt.Println("no destination database specified")
			return
		}
		if dHost.Database != "" {
			dServer.Database = dHost.Database
		}

		// 	Username
		if sHost.Username != "" {
			sServer.Username = sHost.Username
		}
		//
		if dHost.Username != "" {
			dServer.Username = dHost.Username
		}

		// 	Password
		if sHost.Password != "" {
			sServer.Password = sHost.Password
		}
		//
		if dHost.Password != "" {
			dServer.Password = dHost.Password
		}

		// 	Protocol
		if sServer.Protocol == "" && sHost.Protocol != "" {
			sServer.Protocol = sHost.Protocol
		}
		//
		if dServer.Protocol == "" && dHost.Protocol != "" {
			dServer.Protocol = dHost.Protocol
		}

		// 	Schema
		if sServer.Schema == "" && sHost.Schema != "" {
			sServer.Schema = sHost.Schema
		}
		//
		if dServer.Schema == "" && dHost.Schema != "" {
			dServer.Schema = dHost.Schema
		}

		// 	Port
		if sHost.Port != 0 {
			sServer.Port = sHost.Port
		}
		//
		if dHost.Port != 0 {
			dServer.Port = dHost.Port
		}

		// 	Jobcount
		if sHost.Jobcount != 0 {
			sServer.Jobcount = sHost.Jobcount
		}
		//
		if dHost.Jobcount != 0 {
			dServer.Jobcount = dHost.Jobcount
		}

		fmt.Println("sServer", sServer)
		fmt.Println("dServer", dServer)

		fmt.Println("sHost", sHost)
		fmt.Println("dHost", dHost)
		fmt.Println("q", q)
		// fmt.Println("server", server)

		if q.Model == "" {
			fmt.Println("no model specified")
			return
		}
		oSource, err := pkg.OdooConnect(sServer)
		pkg.FatalErr(err)
		oDest, err := pkg.OdooConnect(dServer)
		pkg.FatalErr(err)

		records := pkg.GetRecords(oSource, &q)
		fmt.Println(records)
		umdl := strings.Replace(q.Model, "_", ".", -1)
		for _, record := range records {
			if name, ok := record["name"]; ok {
				fmt.Println(record, name)
				rid, err := oDest.GetID(umdl, oarg{oarg{"id", "=", record["id"]}})
				pkg.CheckErr(err)
				fmt.Println(rid)
				rname, err := oDest.GetID(umdl, oarg{oarg{"name", "=", name}})
				pkg.CheckErr(err)
				fmt.Println(rname)
				if rid == 1 {
					out, err := oDest.Update(umdl, rid, record)
					pkg.CheckErr(err)
					fmt.Println(out)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	copyCmd.Flags().StringVar(&sHost.Hostname, "s_hostname", "", "odoo hostname")
	copyCmd.Flags().StringVar(&sHost.Database, "s_database", "", "odoo database")
	copyCmd.Flags().StringVar(&sHost.Username, "s_username", "", "odoo username")
	copyCmd.Flags().StringVar(&sHost.Password, "s_password", "", "odoo password")
	copyCmd.Flags().StringVar(&sHost.Schema, "s_schema", "", "odoo url schema [http|https]")
	copyCmd.Flags().IntVar(&sHost.Port, "s_port", 0, "odoo port")

	copyCmd.Flags().StringVar(&dHost.Hostname, "d_hostname", "", "odoo hostname")
	copyCmd.Flags().StringVar(&dHost.Database, "d_database", "", "odoo database")
	copyCmd.Flags().StringVar(&dHost.Username, "d_username", "", "odoo username")
	copyCmd.Flags().StringVar(&dHost.Password, "d_password", "", "odoo password")
	copyCmd.Flags().StringVar(&dHost.Schema, "d_schema", "", "odoo url schema [http|https]")
	copyCmd.Flags().IntVar(&dHost.Port, "d_port", 0, "odoo port")

	copyCmd.Flags().StringVarP(&q.Model, "model", "m", "", "model")
	copyCmd.Flags().StringVar(&q.Filter, "filter", "", "filter")
	copyCmd.Flags().IntVar(&q.Offset, "offset", 0, "offset")
	copyCmd.Flags().IntVar(&q.Limit, "limit", 0, "limit")
	copyCmd.Flags().StringVar(&q.Fields, "fields", "", "fields")
	copyCmd.Flags().BoolVar(&q.Count, "count", false, "count")
	copyCmd.MarkFlagRequired("model")
}
