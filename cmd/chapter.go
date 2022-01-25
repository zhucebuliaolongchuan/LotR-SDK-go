/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"LotR-SDK-go/utils"

	"github.com/spf13/cobra"
)

// chapterCmd represents the chapter command
var chapterCmd = &cobra.Command{
	Use:   "chapter",
	Short: "List book chapters",
	Run: func(cmd *cobra.Command, args []string) {
		chapterID, err := cmd.Flags().GetString(utils.ChapterIDCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.ChapterIDCommandFlag, err)
			return
		}

		resp, err := utils.MakeRequest(cmd.Flags(), generateChapterPath(chapterID))
		if err != nil {
			fmt.Printf("could not make the request to the service. %v", err)
			return
		}

		utils.PrintPrettyJson(resp)
	},
}

func init() {
	rootCmd.AddCommand(chapterCmd)

	chapterCmd.Flags().String(utils.URICommandFlag, "https://the-one-api.dev/v2", "specify the API endpoint of the one api service")
	chapterCmd.Flags().String(utils.APIKeyCommandFlag, "", "request data with provided api key")
	chapterCmd.Flags().String(utils.ChapterIDCommandFlag, "", "request one specific book chapter with provided id")
}

func generateChapterPath(chapterID string) string {
	path := "/chapter"

	if chapterID != "" {
		path += fmt.Sprintf("/%s", chapterID)
	}

	return path
}
