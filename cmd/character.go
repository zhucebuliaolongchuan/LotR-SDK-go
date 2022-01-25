/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"LotR-SDK-go/utils"

	"github.com/spf13/cobra"
)

// characterCmd represents the character command
var characterCmd = &cobra.Command{
	Use:   "character",
	Short: "List characters including metadata like name, gender, realm, race and more",
	Run: func(cmd *cobra.Command, args []string) {
		characterID, err := cmd.Flags().GetString(utils.CharacterIDCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.CharacterIDCommandFlag, err)
			return
		}

		quote, err := cmd.Flags().GetBool(utils.GetQuoteCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.GetQuoteCommandFlag, err)
			return
		}

		if quote && characterID == "" {
			fmt.Printf("%s is required if %s is set to true", utils.CharacterIDCommandFlag, utils.GetQuoteCommandFlag)
			return
		}

		resp, err := utils.MakeRequest(cmd.Flags(), generateCharacterPath(characterID, quote))
		if err != nil {
			fmt.Printf("could not make the request to the service. %v", err)
			return
		}

		utils.PrintPrettyJson(resp)
	},
}

func init() {
	rootCmd.AddCommand(characterCmd)

	characterCmd.Flags().String(utils.URICommandFlag, "https://the-one-api.dev/v2", "specify the API endpoint of the one api service")
	characterCmd.Flags().String(utils.APIKeyCommandFlag, "", "request data with provided api key")
	characterCmd.Flags().String(utils.CharacterIDCommandFlag, "", "request one specific character with provided id")
	characterCmd.Flags().Bool(utils.GetQuoteCommandFlag, false, "Request all movie quotes of one specific character if this is set to True")
}

func generateCharacterPath(characterID string, quote bool) string {
	path := "/character"

	if characterID != "" {
		path += fmt.Sprintf("/%s", characterID)
	}

	if quote {
		path += "/quote"
	}

	return path
}
