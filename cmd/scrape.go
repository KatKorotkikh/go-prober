package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Make a request to the target",
	Run: func(cmd *cobra.Command, args []string) {
		targets, _ := cmd.Flags().GetStringArray("target")
		for _, target := range targets {

			result := strings.Split(target, "|")

			if len(result) != 2 {
				log.Fatal("Invalid target should be GET|http://example.com")
			}

			method := result[0]
			url := result[1]
			req, err := http.NewRequest(method, url, nil)

			if err != nil {
				log.Fatal(err)
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				log.Fatalf("Fail to query target: %s", err)
			}

			fmt.Printf("%d %s\n", resp.StatusCode, target)
		}
	},
}

func init() {
	scrapeCmd.Flags().StringArray("target", []string{}, "The target method and url to scrape, like GET|http://example.com")
	scrapeCmd.MarkFlagRequired("target")
}
