package tags

//uses https://api.stackexchange.com/docs
//https://api.stackexchange.com/docs/tags#order=desc&sort=name&inname=java&filter=default&site=stackoverflow&run=true

import (
	"fmt"
	"github.com/laktek/Stack-on-Go/stackongo"
)

func Validate(tags []string) []string {
	session := stackongo.NewSession("stackoverflow")
	params := make(stackongo.Params)
	params.AddVectorized("inname", tags)
	allTags, _ := session.AllTags(params)
	fmt.Printf("Total: %d", allTags.Total)
	for _, tag := range allTags.Items {
		fmt.Printf(tag.Name + "\n")
	}
	return tags
}
