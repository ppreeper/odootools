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

import "github.com/ppreeper/odoojrpc"

func OdooConnect(host Host) (*odoojrpc.Odoo, error) {
	oc := odoojrpc.Odoo{
		Hostname: host.Hostname,
		Port:     host.Port,
		Database: host.Database,
		Username: host.Username,
		Password: host.Password,
		Schema:   host.Schema,
	}
	err := oc.Login()
	if err != nil {
		return nil, err
	}

	return &oc, nil
}
