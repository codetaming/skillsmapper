package tags

//uses https://api.stackexchange.com/docs

import (
	"fmt"
	"github.com/laktek/Stack-on-Go/stackongo"
)

func Validate(tags []string) []string {
	session := stackongo.NewSession("stackoverflow")
	params := make(stackongo.Params)
	params.Add("inname", tags[0])
	allTags, _ := session.AllTags(params)
	fmt.Printf("Total: %d", allTags.Total)
	for _, tag := range allTags.Items {
		fmt.Printf(tag.Name + "\n")
	}
	return tags
}
