package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/adit1o/train4/helpers"
	"github.com/adit1o/train4/models"
	"github.com/spf13/cobra"
)

/// [ ] what i want / flow:
/// user can put a JAV code, example: `EYAN-181`
/// getting metadata from `r18` api
/// showing metadata to user

var searchCmd = &cobra.Command{
	Use		: "search",
	Short	: "Search JAV videos",
	Example	: "javit search URE-013",
	RunE	: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return fmt.Errorf("missing search query")
		}

		helpers.BLine("Searching for: " + args[0])

		resp, err := http.Get("https://next-javit-api.vercel.app/api/r18/bycode?code=" + args[0])
		if err != nil {
			return fmt.Errorf("failed to search: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("failed to read response body: %v\n", err)
		}

		/// json to struct
		var res models.SearchModel
		err = json.Unmarshal(body, &res)
		if err != nil {
			return err
		} 

		/// list actors
		listActors := []string{}
		for _, actor := range res.Actors {
			listActors = append(listActors, actor.CastNames)
		}

		/// output
		fmt.Println("Code			:", res.VideoCode)
		fmt.Println("Title			:", res.Title)
		fmt.Println("Title JA		:", res.TitleJa)
		fmt.Println("Actor			:", strings.Join(listActors, ", "))
		fmt.Println("Genre			:", strings.Join(res.Genre, ", "))
		fmt.Println("Poster			:", res.Poster)
		fmt.Println("Studio			:", res.Studio)
		helpers.PLine()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}