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
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ppreeper/odoojrpc"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	OC       *odoojrpc.Odoo
}

func (app *Application) GetRecords(q *QueryDef) {
	umdl := strings.Replace(q.Model, "_", ".", -1)

	fields := parseFields(q.Fields)
	if q.Count {
		fields = []string{"id"}
	}

	filtp, err := parseFilter(q.Filter)
	CheckErr(err)

	rr, err := app.OC.SearchRead(umdl, filtp, q.Offset, q.Limit, fields)
	FatalErr(err)
	if q.Count {
		fmt.Println("records:", len(rr))
	} else {
		jsonStr, err := json.MarshalIndent(rr, "", "  ")
		CheckErr(err)
		fmt.Println(string(jsonStr))
	}
}

func GetRecords(oc *odoojrpc.Odoo, q *QueryDef) []map[string]any {
	umdl := strings.Replace(q.Model, "_", ".", -1)

	fields := parseFields(q.Fields)
	if q.Count {
		fields = []string{"id"}
	}

	filtp, err := parseFilter(q.Filter)
	CheckErr(err)

	rr, err := oc.SearchRead(umdl, filtp, q.Offset, q.Limit, fields)
	FatalErr(err)
	// if q.Count {
	// 	fmt.Println("records:", len(rr))
	// } else {
	// 	jsonStr, err := json.MarshalIndent(rr, "", "  ")
	// 	CheckErr(err)
	// 	fmt.Println(string(jsonStr))
	// }
	return rr
}
