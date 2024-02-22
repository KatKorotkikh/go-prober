package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Make a request to the target",
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		method, _ := cmd.Flags().GetString("method")

		req, err := http.NewRequest(method, target, nil)

		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatalf("Fail to query target: %s", err)
		}

		fmt.Printf("%d %s\n", resp.StatusCode, target)
	},
}

func init() {
	scrapeCmd.Flags().String("target", "", "The target url to scrape")
	scrapeCmd.MarkFlagRequired("target")
	scrapeCmd.Flags().String("method", "GET", "http method to use")
}
