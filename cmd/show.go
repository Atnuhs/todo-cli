/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"todo-cli/app"
)

func fuga(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	switch len(args) {
	case 0:
		return []string{"foo1", "bar1", "[ぴよ](https://google.com)"}, cobra.ShellCompDirectiveDefault
	case 1:
		return []string{"foo2", "bar2", "baz2"}, cobra.ShellCompDirectiveDefault
	case 2:
		return []string{"foo3", "bar3", "baz3"}, cobra.ShellCompDirectiveDefault
	default:
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}

// showCmd represents the show command
var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show your task as markdown form",
		Run: func(cmd *cobra.Command, args []string) {
			targetPath := filepath.Join(TodoRoot, fmt.Sprintf("%s.md", dateStr))
			f, err := os.Open(targetPath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fmt.Println(app.ShowContent(f))
		},
	}
)

func init() {
	showCmd.AddCommand([]*cobra.Command{
		{
			Use: "completion",
			Run: func(cmd *cobra.Command, args []string) {
				rootCmd.GenFishCompletion(os.Stdout, cobra.EnableCaseInsensitive)
			},
		},
	}...)
	rootCmd.AddCommand(showCmd)
	showCmd.Flags()
	showCmd.ValidArgsFunction = fuga
}
