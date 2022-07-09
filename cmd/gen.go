/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate code from template in current dir",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
		fmt.Println(args)
		fmt.Println(cmd.Flags().GetString("lang"))
		fmt.Println(cmd.Flags().GetString("name"))
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genCmd.Flags().StringP("lang", "l", "cc", "language")
	genCmd.Flags().StringP("name", "n", "", "file name")
}
