/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"LotR-SDK-go/utils"

	"github.com/spf13/cobra"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "List movie quotes or one specific movie quote",
	Run: func(cmd *cobra.Command, args []string) {
		quoteID, err := cmd.Flags().GetString(utils.QuoteIDCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.QuoteIDCommandFlag, err)
			return
		}

		resp, err := utils.MakeRequest(cmd.Flags(), generateQuotePath(quoteID))
		if err != nil {
			fmt.Printf("could not make the request to the service. %v", err)
			return
		}

		utils.PrintPrettyJson(resp)
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	quoteCmd.Flags().String(utils.URICommandFlag, "https://the-one-api.dev/v2", "specify the API endpoint of the-one-api service")
	quoteCmd.Flags().String(utils.APIKeyCommandFlag, "", "search data with provided api key")
	quoteCmd.Flags().String(utils.QuoteIDCommandFlag, "", "request one specific movie quote with provided id")
}

func generateQuotePath(quoteID string) string {
	path := "/quote"
	if quoteID != "" {
		path += fmt.Sprintf("/%s", quoteID)
	}

	return path
}
