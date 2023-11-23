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
package pkg

import (
	"os"

	"gopkg.in/yaml.v3"
)

type QueryDef struct {
	Model  string
	Filter string
	Offset int
	Limit  int
	Fields string
	Count  bool
}

// Conf config structure
type Host struct {
	Hostname string `default:"localhost" json:"hostname"`
	Database string `default:"odoo" json:"database,omitempty"`
	Username string `default:"odoo" json:"username"`
	Password string `default:"odoo" json:"password"`
	Protocol string `default:"jsonrpc" json:"protocol,omitempty"`
	Schema   string `default:"http" json:"schema,omitempty"`
	Port     int    `default:"8069" json:"port,omitempty"`
	Jobcount int    `default:"1" json:"jobcount,omitempty"`
}

func GetConf(configFile string) map[string]Host {
	yamlFile, err := os.ReadFile(configFile)
	CheckErr(err)
	data := make(map[string]Host)
	err = yaml.Unmarshal(yamlFile, data)
	CheckErr(err)
	return data
}
