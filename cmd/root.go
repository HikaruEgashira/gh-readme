/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	gh "github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-readme",
	Short: "README.md を表示します",
	Long:  `README.md を表示します`,
	Run: func(cmd *cobra.Command, args []string) {
		repo := cmd.Flag("repo").Value.String()

		// check if directory exists
		_, err := os.Stat("workspaces/" + repo)
		if err != nil {
			cloneArgs := []string{"repo", "clone", repo, "workspaces/" + repo, "--", "--depth=1"}
			_, _, err = gh.Exec(cloneArgs...)
			if err != nil {
				log.Fatal(err)
			}
		}

		// read README.md
		readme, err := os.ReadFile("workspaces/" + repo + "/README.md")
		if err != nil {
			fmt.Println("README.md not found")
		}
		fmt.Println(string(readme))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("repo", "r", "", "リポジトリ名")
	rootCmd.MarkFlagRequired("repo")
}
