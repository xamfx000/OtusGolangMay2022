/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hw08_envdir_tool/go-envdir/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the root command.
var rootCmd = &cobra.Command{
	Use: "go-envdir /path/to/environment command [args]",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := internal.ReadDir(args[0])
		internal.RunCmd(args[1:], env)
	},
}

func init() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
