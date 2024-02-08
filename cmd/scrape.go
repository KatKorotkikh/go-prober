package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var target *string
var method *string

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Make a request to the target",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.DefaultClient.Get(*target)
		if err != nil {
			log.Fatalf("Fail to query target: %s", err)
		}
		fmt.Printf("%d %s\n", res.StatusCode, *target)
	},
}

func init() {
	target = scrapeCmd.Flags().String("target", "", "The target url to scrape")
	scrapeCmd.MarkFlagRequired("target")
	method = scrapeCmd.Flags().String("method", "GET", "http method to use")
}
