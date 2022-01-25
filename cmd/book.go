/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"LotR-SDK-go/utils"

	"github.com/spf13/cobra"
)

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "List books",
	Run: func(cmd *cobra.Command, args []string) {
		bookID, err := cmd.Flags().GetString(utils.BookIDCommandFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.BookIDCommandFlag, err)
			return
		}

		getChapter, err := cmd.Flags().GetBool(utils.GetChapterCommnadFlag)
		if err != nil {
			fmt.Printf("could not read parameter %s. %v", utils.GetChapterCommnadFlag, err)
			return
		}

		if bookID == "" && getChapter {
			fmt.Printf("%s could not be empty if %s is set to true", utils.BookIDCommandFlag, utils.GetChapterCommnadFlag)
			return
		}

		resp, err := utils.MakeRequest(cmd.Flags(), generateBookPath(bookID, getChapter))
		if err != nil {
			fmt.Printf("could not make the request to the service. %v", err)
			return
		}

		utils.PrintPrettyJson(resp)
	},
}

func init() {
	rootCmd.AddCommand(bookCmd)

	bookCmd.Flags().String(utils.URICommandFlag, "https://the-one-api.dev/v2", "specify the API endpoint of the-one-api service")
	bookCmd.Flags().String(utils.BookIDCommandFlag, "", "request one specific book with the provided id")
	bookCmd.Flags().Bool(utils.GetChapterCommnadFlag, false, "request all chapters of one specific book with provided id")
}

func generateBookPath(bookID string, getChapter bool) string {
	path := "/book"

	if bookID != "" {
		path += fmt.Sprintf("/%s", bookID)
	}

	if getChapter {
		path += "/chapter"
	}

	return path
}
