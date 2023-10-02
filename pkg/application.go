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
