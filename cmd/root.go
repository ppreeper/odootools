// odootools - a cli tool library to access Odoo instances via Json RPC
// Copyright (C) 2023  Peter Preeper

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package cmd

import (
	"os"

	"github.com/ppreeper/odootools/pkg"
	"github.com/spf13/cobra"
)

var (
	configFile string
	sHost      pkg.Host
	dHost      pkg.Host
	q          pkg.QueryDef
	INSERT     = true
	UPDATE     = false
)

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
