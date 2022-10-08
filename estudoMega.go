//
// You many need to press the "Run" button numerous times
//
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

var ctx = context.Background()

func main() {

	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-05-07,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df, err := imports.LoadFromCSV(ctx, strings.NewReader(csvStr), imports.CSVLoadOptions{InferDataTypes: true, NilValue: &[]string{"NA"}[0]})
	if err != nil {
		panic(err)
	}

	fmt.Print(df.Table())
}