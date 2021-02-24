package repo

import (
	"fmt"
	"os"
)

func GetUrl(cArgs []string) string {
	var repo string
	for i, v := range cArgs {
		if v == "--repo" || v == "-r" {
			if len(cArgs) <= i+1 {
				fmt.Println("No Repo Url Provided")
				os.Exit(1)

			} else {
				repo = cArgs[i+1]
			}
		}
	}
	if repo == "" {
		fmt.Println("No Repo Url Provided")
		os.Exit(1)
	}
	return repo
}
