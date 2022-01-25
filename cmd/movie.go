/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"LotR-SDK-go/utils"

	"github.com/spf13/cobra"
)

// movieCmd represents the movie command
var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "List movies",
	Run: func(cmd *cobra.Command, args []string) {
		movieID, err := cmd.Flags().GetString(utils.MovieIDCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.MovieIDCommandFlag, err)
			return
		}

		quote, err := cmd.Flags().GetBool(utils.GetQuoteCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.GetQuoteCommandFlag, err)
			return
		}

		if quote && movieID == "" {
			fmt.Printf("%s is required if %s is set to true", utils.MovieIDCommandFlag, utils.GetQuoteCommandFlag)
			return
		}

		resp, err := utils.MakeRequest(cmd.Flags(), generateMoviePath(movieID, quote))
		if err != nil {
			fmt.Printf("could not make the request to the service. %v", err)
			return
		}

		utils.PrintPrettyJson(resp)
	},
}

func init() {
	rootCmd.AddCommand(movieCmd)

	movieCmd.Flags().String(utils.URICommandFlag, "https://the-one-api.dev/v2", "specify the API endpoint of the-one-api service")
	movieCmd.Flags().String(utils.APIKeyCommandFlag, "", "request data with provided api key")
	movieCmd.Flags().String(utils.MovieIDCommandFlag, "", "request one specific movie with provided id")
	movieCmd.Flags().Bool(utils.GetQuoteCommandFlag, false, "request all movie quotes for one specific movie (only working for the LotR trilogy) if this is set to True")
}

func generateMoviePath(movieID string, quote bool) string {
	path := "/movie"

	if movieID != "" {
		path += fmt.Sprintf("/%s", movieID)
	}

	if quote {
		path += "/quote"
	}

	return path
}
