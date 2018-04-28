package db

import (
	. "fmt"
	"testing"

	"github.com/drgrib/alfred-bear/comp"
)

func TestStringQuery(t *testing.T) {
	var q = `    
	SELECT DISTINCT
		ZUNIQUEIDENTIFIER, ZTITLE 
	FROM 
		ZSFNOTE 
	WHERE 
		ZARCHIVED=0 
		AND ZTRASHED=0 
	ORDER BY 
		ZMODIFICATIONDATE DESC 
	LIMIT 25
`
	path := comp.Expanduser("~/Library/Containers/net.shinyfrog.bear/Data/Documents/Application Data/database.sqlite")
	lite, err := NewLiteDB(path)
	comp.MustBeNil(err)
	results, err := lite.StringQuery(q)
	comp.MustBeNil(err)
	for _, m := range results {
		for k, v := range m {
			Println(k, v)
		}
		Println()
	}
}