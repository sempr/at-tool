/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/sempr/at-tool/pkg/crawler"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := os.Getwd()
		strs := strings.Split(str, "/")
		lastCode := strs[len(strs)-1]
		contestCode := lastCode
		ps, err := crawler.GetTasks(contestCode)
		if err == nil {
			for _, p := range ps {
				dt, err := crawler.GetProblem(p.URL)
				if err != nil {
					log.Panic(err)
				} else {
					crawler.GenDir(p, dt)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
