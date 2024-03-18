package cmd

import "github.com/spf13/cobra"

/// [ ] what i want / flow:
/// 1. get (list) video from a specific file or directory
/// 2. parse metadata from internet for each video
/// 3. apply and rename from metadata to each video
/// 4. download fanArt / poster and save to video folder

var parseCmd = &cobra.Command{
	Use		: "parse",
	Short	: "Parse Metadata your JAV videos from R18",
	Example	: "javit parse <video/path>",
	RunE	: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}